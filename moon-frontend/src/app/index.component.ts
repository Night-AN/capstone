import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterLink, RouterOutlet } from '@angular/router';
import { MatButtonModule } from '@angular/material/button';
import { MatDividerModule } from '@angular/material/divider';
import { MatExpansionModule } from '@angular/material/expansion';
import { MatIconModule } from '@angular/material/icon';
import { MatListModule } from '@angular/material/list';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatToolbarModule } from '@angular/material/toolbar';

interface NavItem {
  path?: string;
  name: string;
  icon: string;
  children?: NavItem[];
}

@Component({
  selector: 'app-index',
  standalone: true,
  imports: [
    CommonModule,
    RouterLink,
    RouterOutlet,
    MatButtonModule,
    MatDividerModule,
    MatExpansionModule,
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

  navItems: NavItem[] = [
    { path: '/dashboard', name: 'Dashboard', icon: 'dashboard' },
    { path: '/users', name: 'Users', icon: 'people' },
    { path: '/organizations', name: 'Organizations', icon: 'business' },
    { path: '/roles', name: 'Roles', icon: 'badge' },
    { path: '/permissions', name: 'Permissions', icon: 'lock' },
    { path: '/resources', name: 'Resources', icon: 'inventory' },
    { path: '/assets', name: 'Assets', icon: 'devices' },
    { path: '/vulnerabilities', name: 'Vulnerabilities', icon: 'bug_report' },
    { path: '/asset-vulnerabilities', name: 'Asset Vulnerabilities', icon: 'link' },
    {
      name: 'Threat Analysis',
      icon: 'psychology',
      children: [
        { path: '/ai/chat', name: 'AI Chat', icon: 'chat' },
        { path: '/ai/model-config', name: 'Model Config', icon: 'settings_applications' },
        { path: '/ai/prompt-template', name: 'Prompt Template', icon: 'description' },
        { path: '/ai/logs', name: 'AI Logs', icon: 'history' }
      ]
    }
  ];
}
