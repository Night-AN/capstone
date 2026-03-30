import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators, FormsModule } from '@angular/forms';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzCardModule } from 'ng-zorro-antd/card';
import { NzFormModule } from 'ng-zorro-antd/form';
import { NzInputModule } from 'ng-zorro-antd/input';
import { NzTableModule } from 'ng-zorro-antd/table';
import { NzModalModule, NzModalService } from 'ng-zorro-antd/modal';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzDatePickerModule } from 'ng-zorro-antd/date-picker';
import { CommonModule } from '@angular/common';
import { UserService } from 'src/app/services/user.service';

interface User {
  id: string;
  nickname: string;
  name: string;
  email: string;
  createdAt: string;
}

enum ModalTitle {
  Create = '创建用户',
  Edit = '编辑用户'
}

@Component({
  selector: 'app-user-page',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    FormsModule,
    NzButtonModule,
    NzCardModule,
    NzFormModule,
    NzInputModule,
    NzTableModule,
    NzModalModule,
    NzDatePickerModule
  ],
  templateUrl: './user-page.html',
  styleUrl: './user-page.scss',
  providers: [NzMessageService]
})
export class UserPage implements OnInit {
  userForm: FormGroup;
  users: User[] = [];

  isEditing = false;
  currentUserId: string | null = null;

  visible = false;
  modalTitle = ModalTitle.Create;
  isLoading = false;
  loading = false;

  // 搜索相关
  searchName = '';
  searchEmail = '';
  dateRange: Date[] = [];

  constructor(
    private fb: FormBuilder,
    private message: NzMessageService,
    private userService: UserService
  ) {
    this.userForm = this.fb.group({
      nickname: ['', [Validators.required]],
      fullname: ['', [Validators.required]],
      email: ['', [Validators.required, Validators.email]]
    });
  }

  ngOnInit(): void {
    this.loadUsers();
  }

  loadUsers(): void {
    this.loading = true;
    this.userService.getList().subscribe(users => {
      this.users = users[0].edges.map((edge: any) => ({
        id: edge.node.id,
        nickname: edge.node.nickname,
        fullname: edge.node.fullname,
        email: edge.node.email,
        createdAt: edge.node.createdat
      }));
      this.loading = false;
    });
  }

  // 搜索用户
  searchUsers(): void {
    this.loading = true;
    // 这里可以根据搜索条件调用后端API
    // 暂时使用前端过滤
    setTimeout(() => {
      this.loading = false;
    }, 500);
  }

  // 重置搜索
  resetSearch(): void {
    this.searchName = '';
    this.searchEmail = '';
    this.dateRange = [];
  }

  createUser(): void {
    this.resetForm();
    this.modalTitle = ModalTitle.Create;
    this.visible = true;
    this.userService.create(this.userForm.value).subscribe(() => {
      this.loadUsers();
      this.message.success('用户创建成功');
    });
  }

  editUser(user: User): void {
    this.isEditing = true;
    this.currentUserId = user.id;
    this.modalTitle = ModalTitle.Edit;
    this.userForm.patchValue({
      nickname: user.nickname,
      name: user.name,
      email: user.email
    });
    this.visible = true;
  }

  deleteUser(id: string): void {
    this.userService.delete(id).subscribe(() => {
      this.loadUsers();
      this.message.success('用户删除成功');
    });
  }

  handleCancel(): void {
    this.visible = false;
    this.resetForm();
  }

  handleOk(): void {
    if (this.userForm.valid) {
      this.isLoading = true;
      setTimeout(() => {
        if (this.isEditing && this.currentUserId) {
          this.updateUser();
        } else {
          this.createUser();
        }
        this.isLoading = false;
        this.visible = false;
      }, 500);
    } else {
      Object.values(this.userForm.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }

  submitForm(): void {
    if (this.userForm.valid) {
      if (this.isEditing && this.currentUserId) {
        this.updateUser();
      } else {
        this.createUser();
      }
    } else {
      Object.values(this.userForm.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }



  updateUser(): void {
    if (this.currentUserId) {
      const index = this.users.findIndex(user => user.id === this.currentUserId);
      if (index !== -1) {
        this.users[index] = {
          ...this.users[index],
          name: this.userForm.value.name,
          email: this.userForm.value.email
        };
        this.message.success('用户更新成功');
        this.resetForm();
      }
    }
  }


  resetForm(): void {
    this.userForm.reset();
    this.isEditing = false;
    this.currentUserId = null;
  }
}
