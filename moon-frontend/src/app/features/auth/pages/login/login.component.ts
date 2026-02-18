import { CommonModule } from '@angular/common';
import { Component, inject } from '@angular/core';
import {
  FormBuilder,
  FormControl,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '@core/services/auth.service';
import { LoginCredentials } from '@models/auth.data';
import { NotificationService } from '@shared/service/notification/notification.service';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatIconModule } from '@angular/material/icon';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
@Component({
  selector: 'app-login',
  imports: [
    CommonModule, 
    ReactiveFormsModule,
    MatButtonModule,
    MatCardModule,
    MatFormFieldModule,
    MatInputModule,
    MatIconModule,
    MatProgressSpinnerModule
  ],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css',
})
export class LoginComponent {
  private router=inject(Router);
  private authService=inject(AuthService);
  private notificationService=inject(NotificationService);

  constructor() {}

  public loginForm = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.email]),
    password: new FormControl('', [Validators.required]),
  });

  public onLogin():void{
    if (this.loginForm.valid) {
      const credentials: LoginCredentials = {
        email: this.loginForm.value.email!,
        password: this.loginForm.value.password!,
      };

      this.authService.login(credentials).subscribe({
        next: () => {
          this.notificationService.success('登录成功');
          this.router.navigate(['/dashboard']);
        },
        error: (error) => {
          this.notificationService.error('登录失败: ' + error.message);
        }
      });
    }
  }

  public goToRegister():void{
    this.router.navigate(['/auth/register'])
  }
}
