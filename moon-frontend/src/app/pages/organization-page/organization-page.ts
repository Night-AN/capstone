import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzCardModule } from 'ng-zorro-antd/card';
import { NzFormModule } from 'ng-zorro-antd/form';
import { NzInputModule } from 'ng-zorro-antd/input';
import { NzTableModule } from 'ng-zorro-antd/table';
import { NzModalModule } from 'ng-zorro-antd/modal';
import { NzMessageService } from 'ng-zorro-antd/message';
import { CommonModule } from '@angular/common';
import { OrganizationService } from 'src/app/services/organization.service';

interface Organization {
  id: string;
  organization_name: string;
  organization_code: string;
  organization_description: string;
  organization_flag: string;
  parent_id: string | null;
  created_at: string;
  updated_at: string | null;
}

enum ModalTitle {
  Create = '创建组织机构',
  Edit = '编辑组织机构'
}

@Component({
  selector: 'app-organization-page',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    NzButtonModule,
    NzCardModule,
    NzFormModule,
    NzInputModule,
    NzTableModule,
    NzModalModule
  ],
  templateUrl: './organization-page.html',
  styleUrl: './organization-page.scss',
  providers: [NzMessageService]
})
export class OrganizationPage implements OnInit {
  organizationForm: FormGroup;
  organizations: Organization[] = [];

  isEditing = false;
  currentOrganizationId: string | null = null;

  visible = false;
  modalTitle = ModalTitle.Create;
  isLoading = false;

  constructor(
    private fb: FormBuilder,
    private message: NzMessageService,
    private organizationService: OrganizationService
  ) {
    this.organizationForm = this.fb.group({
      organization_name: ['', [Validators.required]],
      organization_code: ['', [Validators.required]],
      organization_description: ['', [Validators.required]],
      organization_flag: ['', [Validators.required]],
      parent_id: ['']
    });
  }

  ngOnInit(): void {
    this.loadOrganizations();
  }

  loadOrganizations(): void {
    this.organizationService.getList().subscribe(organizations => {
      this.organizations = organizations[0].edges.map((edge: any) => ({
        id: edge.node.id,
        organization_name: edge.node.organization_name,
        organization_code: edge.node.organization_code,
        organization_description: edge.node.organization_description,
        organization_flag: edge.node.organization_flag,
        parent_id: edge.node.parent_id,
        created_at: edge.node.created_at,
        updated_at: edge.node.updated_at
      }));
    });
  }

  createOrganization(): void {
    this.resetForm();
    this.modalTitle = ModalTitle.Create;
    this.visible = true;
  }

  editOrganization(organization: Organization): void {
    this.isEditing = true;
    this.currentOrganizationId = organization.id;
    this.modalTitle = ModalTitle.Edit;
    this.organizationForm.patchValue({
      organization_name: organization.organization_name,
      organization_code: organization.organization_code,
      organization_description: organization.organization_description,
      organization_flag: organization.organization_flag,
      parent_id: organization.parent_id
    });
    this.visible = true;
  }

  deleteOrganization(id: string): void {
    this.organizationService.delete(id).subscribe(() => {
      this.loadOrganizations();
      this.message.success('组织机构删除成功');
    });
  }

  handleCancel(): void {
    this.visible = false;
    this.resetForm();
  }

  handleOk(): void {
    if (this.organizationForm.valid) {
      this.isLoading = true;
      setTimeout(() => {
        if (this.isEditing && this.currentOrganizationId) {
          this.updateOrganization();
        } else {
          this.createOrganizationSubmit();
        }
        this.isLoading = false;
        this.visible = false;
      }, 500);
    } else {
      Object.values(this.organizationForm.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }

  createOrganizationSubmit(): void {
    this.organizationService.create(this.organizationForm.value).subscribe(() => {
      this.loadOrganizations();
      this.message.success('组织机构创建成功');
      this.resetForm();
    });
  }

  updateOrganization(): void {
    if (this.currentOrganizationId) {
      this.organizationService.update(this.currentOrganizationId, this.organizationForm.value).subscribe(() => {
        this.loadOrganizations();
        this.message.success('组织机构更新成功');
        this.resetForm();
      });
    }
  }

  resetForm(): void {
    this.organizationForm.reset();
    this.isEditing = false;
    this.currentOrganizationId = null;
  }
}