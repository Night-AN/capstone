import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterLink, RouterOutlet } from '@angular/router';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatListModule } from '@angular/material/list';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatToolbarModule } from '@angular/material/toolbar';

@Component({
  selector: 'app-index',
  standalone: true,
  imports: [
    CommonModule,
    RouterLink,
    RouterOutlet,
    MatButtonModule,
    MatIconModule,
    MatListModule,
    MatSidenavModule,
    MatToolbarModule
  ],
  templateUrl: './index.component.html',
  styleUrl: './index.component.scss'
})
export class IndexComponent {
  title = 'Moon Frontend';

  navItems = [
    { path: '/dashboard', name: 'Dashboard', icon: 'dashboard' },
    { path: '/users', name: 'Users', icon: 'people' },
    { path: '/organizations', name: 'Organizations', icon: 'business' },
    { path: '/roles', name: 'Roles', icon: 'badge' },
    { path: '/permissions', name: 'Permissions', icon: 'lock' },
    { path: '/resources', name: 'Resources', icon: 'inventory' }
  ];
}
