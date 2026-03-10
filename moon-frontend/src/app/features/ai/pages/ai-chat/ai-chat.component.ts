import { CommonModule } from '@angular/common';
import { Component, OnInit, inject, signal, ElementRef, ViewChild, AfterViewChecked } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatChipsModule } from '@angular/material/chips';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatListModule } from '@angular/material/list';
import { MatMenuModule } from '@angular/material/menu';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatSelectModule } from '@angular/material/select';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatToolbarModule } from '@angular/material/toolbar';
import { Router } from '@angular/router';
import { AIService } from '@core/services/ai.service';
import { NotificationService } from '@shared/service/notification/notification.service';

interface ChatMessage {
  role: 'user' | 'assistant';
  content: string;
  timestamp: Date;
}

interface Conversation {
  conversation_id: string;
  title: string;
  model_used: string;
  message_count: number;
  created_at: string;
  updated_at: string;
}

@Component({
  selector: 'app-ai-chat',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    MatButtonModule,
    MatCardModule,
    MatChipsModule,
    MatFormFieldModule,
    MatIconModule,
    MatInputModule,
    MatListModule,
    MatMenuModule,
    MatProgressSpinnerModule,
    MatSelectModule,
    MatSidenavModule,
    MatToolbarModule
  ],
  template: `
    <mat-drawer-container class="chat-container">
      <mat-drawer #sidenav mode="side" opened class="conversation-sidenav">
        <div class="sidenav-header">
          <h3>会话历史</h3>
          <button mat-icon-button (click)="startNewChat()" title="新建对话">
            <mat-icon>add</mat-icon>
          </button>
        </div>
        <mat-nav-list>
          @for (conv of conversations(); track conv.conversation_id) {
            <a mat-list-item 
               (click)="loadConversation(conv.conversation_id)"
               [class.active]="conv.conversation_id === currentConversationId()">
              <mat-icon matListItemIcon>chat</mat-icon>
              <span matListItemTitle>{{ conv.title }}</span>
              <button mat-icon-button [matMenuTriggerFor]="menu" (click)="$event.stopPropagation()">
                <mat-icon>more_vert</mat-icon>
              </button>
              <mat-menu #menu="matMenu">
                <button mat-menu-item (click)="deleteConversation(conv.conversation_id)">
                  <mat-icon>delete</mat-icon>
                  <span>删除</span>
                </button>
              </mat-menu>
            </a>
          }
        </mat-nav-list>
        @if (conversations().length === 0) {
          <div class="empty-conversations">
            <mat-icon>forum</mat-icon>
            <p>暂无对话记录</p>
          </div>
        }
      </mat-drawer>

      <mat-drawer-content class="chat-content">
        <mat-toolbar color="primary" class="chat-toolbar">
          @if (isMobile()) {
            <button mat-icon-button (click)="sidenav.toggle()">
              <mat-icon>menu</mat-icon>
            </button>
          }
          <span>AI 智能助手</span>
          <span class="toolbar-spacer"></span>
          @if (currentModel()) {
            <mat-chip>{{ currentModel() }}</mat-chip>
          }
        </mat-toolbar>

        <div class="messages-container" #messagesContainer>
          @if (messages().length === 0) {
            <div class="welcome-message">
              <mat-icon>psychology</mat-icon>
              <h2>您好！我是AI智能助手</h2>
              <p>我可以帮您解答关于资产管理、安全漏洞等方面的问题。</p>
              <p>请在下方输入您的问题开始对话。</p>
              <div class="quick-prompts">
                <button mat-stroked-button (click)="sendQuickPrompt('帮我列出所有服务器类型的资产')">
                  帮我列出所有服务器类型的资产
                </button>
                <button mat-stroked-button (click)="sendQuickPrompt('最近有哪些高危漏洞？')">
                  最近有哪些高危漏洞？
                </button>
                <button mat-stroked-button (click)="sendQuickPrompt('如何提高网络安全？')">
                  如何提高网络安全？
                </button>
              </div>
            </div>
          }

          @for (msg of messages(); track $index) {
            <div class="message" [class.user]="msg.role === 'user'" [class.assistant]="msg.role === 'assistant'">
              <div class="message-avatar">
                <mat-icon>{{ msg.role === 'user' ? 'person' : 'smart_toy' }}</mat-icon>
              </div>
              <div class="message-content">
                <div class="message-bubble">
                  {{ msg.content }}
                </div>
                <span class="message-time">{{ msg.timestamp | date:'HH:mm' }}</span>
              </div>
            </div>
          }

          @if (loading()) {
            <div class="message assistant loading">
              <div class="message-avatar">
                <mat-icon>smart_toy</mat-icon>
              </div>
              <div class="message-content">
                <div class="message-bubble">
                  <mat-spinner diameter="20"></mat-spinner>
                  <span>AI正在思考中...</span>
                </div>
              </div>
            </div>
          }
        </div>

        <div class="input-container">
          <mat-form-field appearance="outline" class="message-input">
            <mat-label>输入您的问题...</mat-label>
            <input matInput [(ngModel)]="userMessage" (keyup.enter)="sendMessage()" [disabled]="loading()">
            <button mat-icon-button matSuffix (click)="sendMessage()" [disabled]="loading() || !userMessage.trim()">
              <mat-icon>send</mat-icon>
            </button>
          </mat-form-field>
        </div>
      </mat-drawer-content>
    </mat-drawer-container>
  `,
  styles: [`
    .chat-container {
      height: calc(100vh - 64px);
    }
    .conversation-sidenav {
      width: 280px;
      background: #f5f5f5;
      border-right: 1px solid #e0e0e0;
    }
    .sidenav-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 16px;
      border-bottom: 1px solid #e0e0e0;
    }
    .sidenav-header h3 {
      margin: 0;
    }
    .empty-conversations {
      text-align: center;
      padding: 32px;
      color: #999;
    }
    .empty-conversations mat-icon {
      font-size: 48px;
      width: 48px;
      height: 48px;
      margin-bottom: 16px;
    }
    .chat-content {
      display: flex;
      flex-direction: column;
      height: 100%;
    }
    .chat-toolbar {
      flex-shrink: 0;
    }
    .toolbar-spacer {
      flex: 1;
    }
    .messages-container {
      flex: 1;
      overflow-y: auto;
      padding: 16px;
      display: flex;
      flex-direction: column;
      gap: 16px;
    }
    .welcome-message {
      text-align: center;
      padding: 48px;
      color: #666;
    }
    .welcome-message mat-icon {
      font-size: 64px;
      width: 64px;
      height: 64px;
      color: #1976d2;
    }
    .welcome-message h2 {
      margin: 16px 0 8px;
    }
    .quick-prompts {
      display: flex;
      flex-wrap: wrap;
      gap: 8px;
      justify-content: center;
      margin-top: 24px;
    }
    .message {
      display: flex;
      gap: 12px;
      max-width: 80%;
    }
    .message.user {
      align-self: flex-end;
      flex-direction: row-reverse;
    }
    .message-avatar {
      width: 40px;
      height: 40px;
      border-radius: 50%;
      display: flex;
      align-items: center;
      justify-content: center;
      flex-shrink: 0;
    }
    .message.user .message-avatar {
      background: #1976d2;
      color: white;
    }
    .message.assistant .message-avatar {
      background: #4caf50;
      color: white;
    }
    .message-content {
      display: flex;
      flex-direction: column;
      gap: 4px;
    }
    .message-bubble {
      padding: 12px 16px;
      border-radius: 12px;
      white-space: pre-wrap;
      word-break: break-word;
    }
    .message.user .message-bubble {
      background: #1976d2;
      color: white;
      border-bottom-right-radius: 4px;
    }
    .message.assistant .message-bubble {
      background: #f5f5f5;
      color: #333;
      border-bottom-left-radius: 4px;
    }
    .message.loading .message-bubble {
      display: flex;
      align-items: center;
      gap: 8px;
    }
    .message-time {
      font-size: 12px;
      color: #999;
    }
    .input-container {
      padding: 16px;
      background: white;
      border-top: 1px solid #e0e0e0;
    }
    .message-input {
      width: 100%;
    }
    mat-nav-list a.active {
      background: rgba(0,0,0,0.04);
    }
    @media (max-width: 768px) {
      .conversation-sidenav {
        width: 100%;
      }
      .message {
        max-width: 95%;
      }
    }
  `]
})
export class AIChatComponent implements OnInit, AfterViewChecked {
  @ViewChild('messagesContainer') messagesContainer!: ElementRef;
  @ViewChild('sidenav') sidenav: any;

  private aiService = inject(AIService);
  private notificationService = inject(NotificationService);
  private router = inject(Router);

  userMessage = '';
  messages = signal<ChatMessage[]>([]);
  conversations = signal<Conversation[]>([]);
  currentConversationId = signal<string | null>(null);
  currentModel = signal<string>('');
  loading = signal<boolean>(false);

  ngOnInit(): void {
    this.loadConversations();
  }

  ngAfterViewChecked(): void {
    this.scrollToBottom();
  }

  isMobile(): boolean {
    return window.innerWidth < 768;
  }

  scrollToBottom(): void {
    if (this.messagesContainer) {
      const container = this.messagesContainer.nativeElement;
      container.scrollTop = container.scrollHeight;
    }
  }

  loadConversations(): void {
    this.aiService.listConversations().subscribe({
      next: (response) => {
        this.conversations.set(response.data?.conversations || []);
      },
      error: (error) => {
        console.error('Error loading conversations:', error);
      }
    });
  }

  loadConversation(conversationId: string): void {
    this.currentConversationId.set(conversationId);
    this.aiService.getConversation(conversationId).subscribe({
      next: (response) => {
        const chatMessages = response.data?.messages || [];
        const loadedMessages: ChatMessage[] = chatMessages.map((m: any) => ({
          role: m.role,
          content: m.content,
          timestamp: new Date(m.created_at)
        }));
        this.messages.set(loadedMessages);
        this.currentModel.set(response.data?.model_used || '');
        this.scrollToBottom();
      },
      error: (error) => {
        this.notificationService.error('加载对话失败');
      }
    });
  }

  startNewChat(): void {
    this.currentConversationId.set(null);
    this.messages.set([]);
    this.currentModel.set('');
    this.userMessage = '';
  }

  deleteConversation(conversationId: string): void {
    if (confirm('确定要删除这个对话吗？')) {
      this.aiService.deleteConversation(conversationId).subscribe({
        next: () => {
          this.notificationService.success('删除成功');
          if (this.currentConversationId() === conversationId) {
            this.startNewChat();
          }
          this.loadConversations();
        },
        error: (error) => {
          this.notificationService.error('删除失败');
        }
      });
    }
  }

  sendQuickPrompt(prompt: string): void {
    this.userMessage = prompt;
    this.sendMessage();
  }

  sendMessage(): void {
    if (!this.userMessage.trim() || this.loading()) return;

    const userMsg: ChatMessage = {
      role: 'user',
      content: this.userMessage,
      timestamp: new Date()
    };
    this.messages.update(msgs => [...msgs, userMsg]);

    const messageToSend = this.userMessage;
    this.userMessage = '';
    this.loading.set(true);

    this.aiService.chat(
      messageToSend,
      this.currentConversationId() || undefined
    ).subscribe({
      next: (response) => {
        this.loading.set(false);
        
        const assistantMsg: ChatMessage = {
          role: 'assistant',
          content: response.data?.response || '抱歉，我没有收到有效的回复。',
          timestamp: new Date()
        };
        this.messages.update(msgs => [...msgs, assistantMsg]);

        if (response.data?.conversation_id) {
          this.currentConversationId.set(response.data.conversation_id);
        }
        if (response.data?.model_used) {
          this.currentModel.set(response.data.model_used);
        }

        this.loadConversations();
        this.scrollToBottom();
      },
      error: (error) => {
        this.loading.set(false);
        this.notificationService.error('发送消息失败，请稍后重试');
        console.error('Chat error:', error);
      }
    });
  }
}
