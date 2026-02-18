import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-test',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './test.component.html',
  styleUrls: ['./test.component.scss']
})
export class TestComponent {
  directoryStructure = {
    features: {
      auth: {
        pages: {
          login: ['login.component.css', 'login.component.html', 'login.component.spec.ts', 'login.component.ts'],
          register: ['register.component.css', 'register.component.html', 'register.component.spec.ts', 'register.component.ts']
        },
        files: ['auth-routing.module.ts', 'auth.module.ts']
      },
      dashboard: {
        files: ['dashboard.component.css', 'dashboard.component.html', 'dashboard.component.ts']
      },
      organization: {
        pages: {
          'organization-detail': ['organization-detail.component.html', 'organization-detail.component.scss', 'organization-detail.component.ts'],
          'organization-form': ['organization-form.component.html', 'organization-form.component.scss', 'organization-form.component.ts'],
          'organization-management': ['organization-management.component.html', 'organization-management.component.scss', 'organization-management.component.ts']
        },
        files: ['organization-routing.module.ts', 'organization.module.ts']
      },
      permission: {
        pages: {
          'permission-detail': ['permission-detail.component.html', 'permission-detail.component.scss', 'permission-detail.component.ts'],
          'permission-form': ['permission-form.component.html', 'permission-form.component.scss', 'permission-form.component.ts'],
          'permission-management': ['permission-management.component.html', 'permission-management.component.scss', 'permission-management.component.ts']
        },
        files: ['permission-routing.module.ts', 'permission.module.ts']
      },
      resource: {
        pages: {
          'resource-detail': ['resource-detail.component.html', 'resource-detail.component.scss', 'resource-detail.component.ts'],
          'resource-form': ['resource-form.component.html', 'resource-form.component.scss', 'resource-form.component.ts'],
          'resource-management': ['resource-management.component.html', 'resource-management.component.scss', 'resource-management.component.ts']
        },
        files: ['resource-routing.module.ts', 'resource.module.ts']
      },
      role: {
        pages: {
          'role-detail': ['role-detail.component.html', 'role-detail.component.scss', 'role-detail.component.ts'],
          'role-form': ['role-form.component.html', 'role-form.component.scss', 'role-form.component.ts'],
          'role-management': ['role-management.component.html', 'role-management.component.scss', 'role-management.component.ts']
        },
        files: ['role-routing.module.ts', 'role.module.ts']
      },
      user: {
        pages: {
          'user-detail': ['user-detail.component.html', 'user-detail.component.scss', 'user-detail.component.ts'],
          'user-form': ['user-form.component.html', 'user-form.component.scss', 'user-form.component.ts'],
          'user-management': ['user-management.component.html', 'user-management.component.scss', 'user-management.component.ts']
        },
        files: ['user-routing.module.ts', 'user.module.ts']
      }
    }
  };
}
