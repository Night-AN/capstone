import { Component, OnInit, inject, signal } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormBuilder, FormGroup, Validators, ReactiveFormsModule } from '@angular/forms';
import { RoleService } from '@core/services/role.service';
import { NotificationService } from '@shared/service/notification/notification.service';
import { Role, RoleCreateRequest, RoleUpdateRequest } from '@models/role.model';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-role-form',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatIconModule,
    MatProgressSpinnerModule
  ],
  templateUrl: './role-form.component.html',
  styleUrl: './role-form.component.scss'
})
export class RoleFormComponent implements OnInit {
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private fb = inject(FormBuilder);
  private roleService = inject(RoleService);
  private notificationService = inject(NotificationService);

  roleId: string | null = null;
  roleForm: FormGroup;
  loading = signal<boolean>(false);
  error = signal<string | null>(null);

  constructor() {
    this.roleForm = this.fb.group({
      role_name: ['', Validators.required],
      description: [''],
      sensitive_flag: [false]
    });
  }

  ngOnInit(): void {
    this.roleId = this.route.snapshot.paramMap.get('id');
    
    if (this.roleId) {
      this.loadRoleData();
    }
  }

  get isEditMode(): boolean {
    return !!this.roleId;
  }

  loadRoleData(): void {
    if (!this.roleId) return;

    this.loading.set(true);
    this.error.set(null);

    this.roleService.getRoleById(this.roleId).subscribe({
      next: (role) => {
        if (role) {
          this.roleForm.patchValue({
            role_name: role.role_name,
            description: role.description,
            sensitive_flag: role.sensitive_flag
          });
        }
        this.loading.set(false);
      },
      error: (err) => {
        console.error('Error loading role data:', err);
        this.error.set('Failed to load role data');
        this.loading.set(false);
        this.notificationService.error('Failed to load role data');
      }
    });
  }

  onSubmit(): void {
    if (this.roleForm.invalid) {
      this.roleForm.markAllAsTouched();
      return;
    }

    this.loading.set(true);
    this.error.set(null);

    const formData = this.roleForm.value;
    
    if (this.isEditMode && this.roleId) {
      // Update role
      const updateRequest: RoleUpdateRequest = {
        role_id: this.roleId,
        role_name: formData.role_name,
        description: formData.description,
        sensitive_flag: formData.sensitive_flag
      };

      this.roleService.updateRole(updateRequest).subscribe({
        next: (role) => {
          this.notificationService.success('Role updated successfully');
          this.router.navigate(['/role/detail', role.role_id]);
        },
        error: (err) => {
          this.error.set('Failed to update role');
          this.loading.set(false);
          this.notificationService.error('Failed to update role');
        }
      });
    } else {
      // Create role
      const createRequest: RoleCreateRequest = {
        role_name: formData.role_name,
        description: formData.description,
        sensitive_flag: formData.sensitive_flag
      };

      this.roleService.createRole(createRequest).subscribe({
        next: (role) => {
          this.notificationService.success('Role created successfully');
          this.router.navigate(['/role/detail', role.role_id]);
        },
        error: (err) => {
          this.error.set('Failed to create role');
          this.loading.set(false);
          this.notificationService.error('Failed to create role');
        }
      });
    }
  }

  onCancel(): void {
    if (this.isEditMode && this.roleId) {
      this.router.navigate(['/role/detail', this.roleId]);
    } else {
      this.router.navigate(['/role']);
    }
  }
}