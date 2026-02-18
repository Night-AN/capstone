import { CommonModule } from '@angular/common';
import { Component, OnInit, ViewChild, inject } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatPaginator, MatPaginatorModule } from '@angular/material/paginator';
import { MatSort, MatSortModule } from '@angular/material/sort';
import { MatTableDataSource, MatTableModule } from '@angular/material/table';
import { Router } from '@angular/router';
import { OrganizationService } from '@core/services/organization.service';
import { OrganizationListItem } from '@models/organization.model';
import { NotificationService } from '@shared/service/notification/notification.service';

@Component({
  selector: 'app-organization-management',
  imports: [
    CommonModule,
    FormsModule,
    MatTableModule,
    MatPaginatorModule,
    MatSortModule,
    MatButtonModule,
    MatIconModule,
    MatFormFieldModule,
    MatInputModule
  ],
  templateUrl: './organization-management.component.html',
  styleUrl: './organization-management.component.scss'
})
export class OrganizationManagementComponent implements OnInit {
  displayedColumns: string[] = ['name', 'description', 'created_at', 'actions'];
  dataSource = new MatTableDataSource<OrganizationListItem>();
  searchKeyword: string = '';

  @ViewChild(MatPaginator) paginator!: MatPaginator;
  @ViewChild(MatSort) sort!: MatSort;

  private organizationService = inject(OrganizationService);
  private router = inject(Router);
  private notificationService = inject(NotificationService);

  constructor() { }

  ngOnInit(): void {
    this.loadOrganizations();
  }

  ngAfterViewInit() {
    this.dataSource.paginator = this.paginator;
    this.dataSource.sort = this.sort;
  }

  loadOrganizations(): void {
    this.organizationService.getOrganizations().subscribe({
      next: (response: any) => {
        this.dataSource.data = response.data?.organizations || [];
      },
      error: (error: any) => {
        this.notificationService.error('加载组织失败');
      }
    });
  }

  applyFilter(): void {
    this.dataSource.filter = this.searchKeyword.trim().toLowerCase();
  }

  createOrganization(): void {
    this.router.navigate(['/organizations/create']);
  }

  editOrganization(organizationId: string): void {
    this.router.navigate(['/organizations/edit', organizationId]);
  }

  viewOrganization(organizationId: string): void {
    this.router.navigate(['/organizations/detail', organizationId]);
  }

  deleteOrganization(organizationId: string): void {
    if (confirm('确定要删除这个组织吗？')) {
      this.organizationService.deleteOrganization({ organization_id: organizationId }).subscribe({
        next: (response: any) => {
          this.notificationService.success('删除组织成功');
          this.loadOrganizations();
        },
        error: (error: any) => {
          this.notificationService.error('删除组织失败');
        }
      });
    }
  }
}
