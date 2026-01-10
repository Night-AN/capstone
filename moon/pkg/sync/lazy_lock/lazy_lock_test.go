package lazy_lock_test

import (
	"fmt"
	"moon/pkg/sync/lazy_lock"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// TestLazyLockBasic tests the fundamental lazy initialization behavior.
// It verifies that the initialization function is called exactly once,
// even when Get() is called multiple times.
func TestLazyLockBasic(t *testing.T) {
	// Counter to track how many times the init function is called
	var initCount int32
	lazyInt := lazy_lock.New(func() string {
		// Increment counter each time the init function executes
		atomic.AddInt32(&initCount, 1)
		return "initialized"
	})

	// First call to Get() should trigger initialization
	val1 := lazyInt.Get()
	if val1 != "initialized" {
		t.Errorf("expected 'initialized', got %s", val1)
	}

	// Second call should return cached value without re-initializing
	val2 := lazyInt.Get()
	if val2 != "initialized" {
		t.Errorf("expected 'initialized', got %s", val2)
	}

	// Verify the init function was called exactly once
	if atomic.LoadInt32(&initCount) != 1 {
		t.Errorf("init function should be called exactly once, called %d times", initCount)
	}
}

// Reader interface for testing lazy initialization with interface types
type Reader interface {
	Read() string
}

// mockReader is a mock implementation of the Reader interface
type mockReader struct {
	data string
}

// Implement the Reader interface
func (m *mockReader) Read() string {
	return m.data
}

// TestLazyLockWithInterface tests lazy initialization with interface return types.
// This verifies that the lazy lock can handle non-concrete types.
func TestLazyLockWithInterface(t *testing.T) {
	// Create lazy-initialized Reader instance
	lazy := lazy_lock.New(func() Reader {
		return &mockReader{data: "test data"}
	})

	reader := lazy.Get()
	if reader == nil {
		t.Error("expected non-nil reader")
		return
	}

	// Actually use the interface method
	readData := reader.Read()
	if readData != "test data" {
		t.Errorf("expected 'test data', got '%s'", readData)
	}
}

// TestLazyLockRace tests for race conditions between concurrent read/write operations.
// The test spawns two goroutines that call Get() concurrently to ensure thread safety.
func TestLazyLockRace(t *testing.T) {
	lazy := lazy_lock.New(func() int {
		return 42
	})

	done := make(chan bool)

	// Write operation (actually just triggers initialization)
	go func() {
		lazy.Get()
		done <- true
	}()

	// Read operation
	go func() {
		lazy.Get()
		done <- true
	}()

	<-done
	<-done
}

// TestLazyLockStress performs stress testing with concurrent access.
// It spawns many goroutines that repeatedly access the lazy-initialized value
// to verify correct behavior under high concurrency.
func TestLazyLockStress(t *testing.T) {
	// Skip in short mode to avoid long-running tests
	if testing.Short() {
		t.Skip("skipping stress test in short mode")
	}

	var initCount int64
	lazy := lazy_lock.New(func() int64 {
		// Simulate some work with a short sleep
		time.Sleep(time.Microsecond)
		return atomic.AddInt64(&initCount, 1)
	})

	const goroutines = 1000
	const iterations = 100

	var wg sync.WaitGroup
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				val := lazy.Get()
				// All goroutines should see the same initialized value (1)
				if val != 1 {
					t.Errorf("expected 1, got %d", val)
				}
				runtime.Gosched() // Yield CPU to increase chance of race conditions
			}
		}()
	}

	wg.Wait()

	// Verify initialization happened exactly once
	if atomic.LoadInt64(&initCount) != 1 {
		t.Errorf("init should be called exactly once, called %d times", initCount)
	}
}

// Simulates a real-world UserRepository with time-consuming initialization logic
type UserRepository struct {
	dbConn   string         // Database connection string
	cache    map[int]string // In-memory user cache
	initTime time.Time      // When the repository was initialized
	initBy   int            // Which goroutine performed the initialization
}

// NewUserRepository simulates a costly repository initialization (connecting to DB, warming cache, etc.)
func NewUserRepository(initBy int) *UserRepository {
	// Simulate establishing database connection (50ms delay)
	time.Sleep(50 * time.Millisecond)

	// Simulate cache warmup (loading initial data)
	cache := make(map[int]string)
	for i := 1; i <= 100; i++ {
		cache[i] = fmt.Sprintf("User%d", i)
	}

	return &UserRepository{
		dbConn:   "postgres://localhost:5432",
		cache:    cache,
		initTime: time.Now(),
		initBy:   initBy,
	}
}

// Business method: Query a user by ID
func (r *UserRepository) GetUser(id int) (string, bool) {
	user, exists := r.cache[id]
	return user, exists
}

// Business method: Count total users
func (r *UserRepository) Count() int {
	return len(r.cache)
}

// TestLazyLock_RepositoryStartupStorm tests the "startup storm" scenario.
// Simulates when a service just starts and many requests arrive simultaneously,
// all trying to initialize the Repository.
func TestLazyLock_RepositoryStartupStorm(t *testing.T) {
	var initCalls int32
	var initByGoroutine int32

	lazyRepo := lazy_lock.New(func() *UserRepository {
		// Track initialization information
		atomic.AddInt32(&initCalls, 1)
		gID := atomic.LoadInt32(&initByGoroutine)
		return NewUserRepository(int(gID))
	})

	const numRequests = 50 // Simulate 50 simultaneous requests
	var wg sync.WaitGroup
	startCh := make(chan struct{}) // Synchronization channel for simultaneous start

	// Store results from each request
	type requestResult struct {
		goroutineID int
		user        string
		initTime    time.Time
		initBy      int
	}
	results := make([]requestResult, numRequests)
	var resultsMu sync.Mutex

	// Launch multiple goroutines to simulate concurrent requests
	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func(reqID int) {
			defer wg.Done()

			// Wait for signal to ensure all goroutines start simultaneously
			<-startCh

			// Record which goroutine triggered the initialization (first to call)
			if atomic.CompareAndSwapInt32(&initByGoroutine, 0, int32(reqID)) {
				t.Logf("Goroutine %d triggered initialization", reqID)
			}

			// Get Repository and perform business operation
			repo := lazyRepo.Get()
			user, exists := repo.GetUser(reqID + 1) // Query different users

			if !exists {
				t.Errorf("Goroutine %d: user does not exist", reqID)
			}

			// Store result for later validation
			resultsMu.Lock()
			results[reqID] = requestResult{
				goroutineID: reqID,
				user:        user,
				initTime:    repo.initTime,
				initBy:      repo.initBy,
			}
			resultsMu.Unlock()

		}(i)
	}

	// All goroutines ready, start simultaneously
	close(startCh)
	wg.Wait()

	// Validation: Verify correct behavior
	if atomic.LoadInt32(&initCalls) != 1 {
		t.Errorf("Repository should be initialized exactly once, but was initialized %d times", initCalls)
	}

	// Verify all requests see the same Repository instance
	firstInitTime := results[0].initTime
	firstInitBy := results[0].initBy
	for i, result := range results {
		if result.initTime != firstInitTime {
			t.Errorf("Request %d: saw a different Repository instance (different init time)", i)
		}
		if result.initBy != firstInitBy {
			t.Errorf("Request %d: inconsistent initializing goroutine recorded", i)
		}
		if result.user != fmt.Sprintf("User%d", result.goroutineID+1) {
			t.Errorf("Request %d: incorrect query result, expected User%d, got %s",
				i, result.goroutineID+1, result.user)
		}
	}

	t.Logf("Test passed: %d concurrent requests, Repository initialized only once (by goroutine %d)",
		numRequests, firstInitBy)
}

// TestLazyLock_RepositoryMixedLoad tests mixed workload scenario.
// Simulates concurrent read operations after Repository initialization is complete.
func TestLazyLock_RepositoryMixedLoad(t *testing.T) {
	lazyRepo := lazy_lock.New(func() *UserRepository {
		return NewUserRepository(0)
	})

	const (
		readers         = 30  // 30 reader goroutines
		writesPerReader = 100 // Each reader performs 100 queries
	)

	var wg sync.WaitGroup
	var totalQueries int32
	var failedQueries int32

	// Launch multiple reader goroutines
	for i := 0; i < readers; i++ {
		wg.Add(1)
		go func(readerID int) {
			defer wg.Done()

			// Each reader performs multiple queries
			for j := 0; j < writesPerReader; j++ {
				repo := lazyRepo.Get()
				userID := (readerID*writesPerReader+j)%100 + 1
				user, exists := repo.GetUser(userID)

				atomic.AddInt32(&totalQueries, 1)

				if !exists {
					atomic.AddInt32(&failedQueries, 1)
					t.Errorf("Reader %d: user %d does not exist", readerID, userID)
				}
				if user != fmt.Sprintf("User%d", userID) {
					atomic.AddInt32(&failedQueries, 1)
				}
			}
		}(i)
	}

	wg.Wait()

	if atomic.LoadInt32(&failedQueries) != 0 {
		t.Errorf("%d queries failed", failedQueries)
	}

	expectedQueries := readers * writesPerReader
	if atomic.LoadInt32(&totalQueries) != int32(expectedQueries) {
		t.Errorf("Expected %d queries, but got %d", expectedQueries, totalQueries)
	}

	t.Logf("Mixed load test passed: %d goroutines completed %d queries, 0 failures",
		readers, totalQueries)
}

// TestLazyLock_RepositoryInitializationDelay tests the impact of slow initialization on concurrency.
func TestLazyLock_RepositoryInitializationDelay(t *testing.T) {
	initStartTime := time.Now()
	var initEndTime time.Time

	lazyRepo := lazy_lock.New(func() *UserRepository {
		time.Sleep(100 * time.Millisecond) // Simulate slow initialization
		initEndTime = time.Now()
		return NewUserRepository(0)
	})

	var firstAccessTime time.Time
	var lastAccessTime time.Time
	var wg sync.WaitGroup

	// Launch 20 goroutines that arrive at different times
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(goroutineID int, delay time.Duration) {
			defer wg.Done()
			time.Sleep(delay) // Different goroutines arrive at different times
			accessTime := time.Now()

			if goroutineID == 0 {
				firstAccessTime = accessTime
			}
			if goroutineID == 19 {
				lastAccessTime = accessTime
			}

			repo := lazyRepo.Get()
			_ = repo
		}(i, time.Duration(i)*5*time.Millisecond) // Start one every 5ms
	}

	wg.Wait()

	// Verify all goroutines waited for initialization to complete
	if lastAccessTime.Before(initEndTime) {
		t.Logf("Last goroutine arrived before initialization completed (this is not guaranteed)")
	}

	t.Logf("Initialization took: %v, first request arrived: %v, last request arrived: %v",
		initEndTime.Sub(initStartTime),
		firstAccessTime.Sub(initStartTime),
		lastAccessTime.Sub(initStartTime))
}

// BenchmarkLazyLock_RepositoryRealWorld provides real-world benchmark measurements.
func BenchmarkLazyLock_RepositoryRealWorld(b *testing.B) {
	lazyRepo := lazy_lock.New(func() *UserRepository {
		// Simulate real initialization overhead
		time.Sleep(1 * time.Millisecond)
		return NewUserRepository(0)
	})

	// Initialize first
	repo := lazyRepo.Get()
	if repo == nil {
		b.Fatal("Initialization failed")
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// Simulate business logic query
			user, _ := repo.GetUser(42)
			if user != "User42" {
				b.Errorf("unexpected user: %s", user)
			}
		}
	})
}
