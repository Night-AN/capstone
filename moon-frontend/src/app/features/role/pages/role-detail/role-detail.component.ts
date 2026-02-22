import { Component, OnInit, inject, signal } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { RoleService } from '@core/services/role.service';
import { NotificationService } from '@shared/service/notification/notification.service';
import { Role } from '@models/role.model';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { MatTableModule } from '@angular/material/table';

@Component({
  selector: 'app-role-detail',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    MatButtonModule,
    MatIconModule,
    MatProgressSpinnerModule,
    MatTableModule
  ],
  templateUrl: './role-detail.component.html',
  styleUrl: './role-detail.component.scss'
})
export class RoleDetailComponent implements OnInit {
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private roleService = inject(RoleService);
  private notificationService = inject(NotificationService);

  roleId: string | null = null;
  role = signal<Role | null>(null);
  loading = signal<boolean>(false);
  error = signal<string | null>(null);

  ngOnInit(): void {
    this.roleId = this.route.snapshot.paramMap.get('id');
    if (this.roleId) {
      this.loadRoleDetail();
    }
  }

  loadRoleDetail(): void {
    if (!this.roleId) return;

    this.loading.set(true);
    this.error.set(null);

    this.roleService.getRoleById(this.roleId).subscribe({
      next: (role) => {
        this.role.set({
          role_id: role.role_id,
          role_name: role.role_name,
          description: role.description,
          sensitive_flag: role.sensitive_flag,
          created_at: role.created_at,
          updated_at: role.updated_at
        });
        this.loading.set(false);
      },
      error: (err) => {
        this.error.set('Failed to load role details');
        this.loading.set(false);
        this.notificationService.error('Failed to load role details');
      }
    });
  }

  editRole(): void {
    if (this.roleId) {
      this.router.navigate(['/role/edit', this.roleId]);
    }
  }

  deleteRole(): void {
    if (!this.roleId) return;

    if (confirm('Are you sure you want to delete this role?')) {
      this.roleService.deleteRole({ role_id: this.roleId }).subscribe({
        next: () => {
          this.notificationService.success('Role deleted successfully');
          this.router.navigate(['/role']);
        },
        error: (err) => {
          this.notificationService.error('Failed to delete role');
        }
      });
    }
  }

  backToList(): void {
    this.router.navigate(['/role']);
  }
}