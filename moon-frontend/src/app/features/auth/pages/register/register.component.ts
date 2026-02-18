import { CommonModule } from '@angular/common';
import { Component, inject, signal } from '@angular/core';
import {
  AbstractControl,
  FormBuilder,
  FormControl,
  FormGroup,
  ReactiveFormsModule,
  ValidationErrors,
  ValidatorFn,
  Validators,
} from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '@core/services/auth.service';
import { RegisterData } from '@models/auth.data';
import { NotificationService } from '@shared/service/notification/notification.service';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatIconModule } from '@angular/material/icon';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';

export function passwordMatchValidator(): ValidatorFn {
  return (control: AbstractControl): ValidationErrors | null => {
    const password = control.get('password');
    const confirmPassword = control.get('confirmPassword');

    return password?.value ===confirmPassword?.value ? null :{
      passwordMismatch :true
    }
  };
}

@Component({
  selector: 'app-register',
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
  templateUrl: './register.component.html',
  styleUrl: './register.component.css',
})
export class RegisterComponent {
  private router = inject(Router);
  private formBuilder = inject(FormBuilder);
  private authService=inject(AuthService);
  private notificationService=inject(NotificationService);
  
  public showPasswordError = signal(false)

  constructor() {}

  public registerFrom = this.formBuilder.group({
    nickname: ['', [Validators.required, Validators.minLength(3)]],
    fullName: ['', [Validators.required, Validators.minLength(3)]],
    email: ['', [Validators.required, Validators.email]],
    password: ['', [Validators.required,Validators.minLength(6)]],
    confirmPassword: ['', [Validators.required,Validators.minLength(6)]],
  }, {
    validators: passwordMatchValidator()
  });


  public onRegister(): void {
    if (this.registerFrom.valid) {
      const registerData: RegisterData = {
        nickname: this.registerFrom.value.nickname!,
        fullName: this.registerFrom.value.fullName!,
        email: this.registerFrom.value.email!,
        password: this.registerFrom.value.password!,
      };

      this.authService.register(registerData).subscribe({
        next: () => {
          this.notificationService.success('注册成功');
          this.router.navigate(['/auth/login']);
        },
        error: (error) => {
          this.notificationService.error('注册失败: ' + error.message);
        }
      });
    }else{
      this.showPasswordError.set(true)
    }
  }

  public onPasswordInput():void{
    this.showPasswordError.set(false)
  }

  public goToLogin(): void {
    this.router.navigate(['/auth/login']);
  }
}
