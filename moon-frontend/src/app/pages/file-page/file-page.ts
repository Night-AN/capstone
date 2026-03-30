import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators, ReactiveFormsModule, FormsModule } from '@angular/forms';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzModalService } from 'ng-zorro-antd/modal';
import { NzTableModule } from 'ng-zorro-antd/table';
import { NzPaginationModule } from 'ng-zorro-antd/pagination';
import { NzInputModule } from 'ng-zorro-antd/input';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzCardModule } from 'ng-zorro-antd/card';
import { NzFormModule } from 'ng-zorro-antd/form';
import { NzModalModule } from 'ng-zorro-antd/modal';
import { CommonModule } from '@angular/common';
import { FileService } from '../../services/file.service';

@Component({
  selector: 'app-file-page',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    FormsModule,
    NzTableModule,
    NzPaginationModule,
    NzInputModule,
    NzButtonModule,
    NzIconModule,
    NzCardModule,
    NzFormModule,
    NzModalModule
  ],
  templateUrl: './file-page.html',
  styleUrls: ['./file-page.less']
})
export class FilePageComponent implements OnInit {
  // 搜索相关
  searchValue: string = '';
  
  // 分页相关
  currentPage: number = 1;
  pageSize: number = 10;
  totalCount: number = 0;
  
  // 数据相关
  fileList: any[] = [];
  loading: boolean = false;
  
  // 弹窗相关
  createModalVisible: boolean = false;
  updateModalVisible: boolean = false;
  currentFile: any = null;
  
  // 表单相关
  createForm: FormGroup;
  updateForm: FormGroup;

  constructor(
    private fb: FormBuilder,
    private fileService: FileService,
    private message: NzMessageService,
    private modal: NzModalService
  ) {
    // 初始化表单
    this.createForm = this.fb.group({
      fileName: ['', Validators.required],
      fileType: ['', Validators.required],
      fileSize: [0, [Validators.required, Validators.min(0)]]
    });
    
    this.updateForm = this.fb.group({
      fileName: ['', Validators.required],
      fileType: ['', Validators.required],
      fileSize: [0, [Validators.required, Validators.min(0)]]
    });
  }

  ngOnInit(): void {
    this.getFiles();
  }

  // 获取文件列表
  getFiles() {
    this.loading = true;
    const where = this.searchValue ? { fileName: { contains: this.searchValue } } : {};
    
    this.fileService.getList(where, null, this.pageSize, ((this.currentPage - 1) * this.pageSize).toString()).subscribe({
      next: (response) => {
        this.fileList = response.data.files.edges.map((edge: any) => edge.node);
        this.totalCount = response.data.files.totalCount;
        this.loading = false;
      },
      error: (error) => {
        console.error('获取文件列表失败:', error);
        this.message.error('获取文件列表失败');
        this.loading = false;
      }
    });
  }

  // 搜索文件
  searchFiles() {
    this.currentPage = 1;
    this.getFiles();
  }

  // 分页变化
  pageIndexChange(page: number) {
    this.currentPage = page;
    this.getFiles();
  }

  pageSizeChange(size: number) {
    this.pageSize = size;
    this.currentPage = 1;
    this.getFiles();
  }

  // 打开创建弹窗
  openCreateModal() {
    this.createForm.reset();
    this.createModalVisible = true;
  }

  // 打开更新弹窗
  openUpdateModal(file: any) {
    this.currentFile = file;
    this.updateForm.patchValue({
      fileName: file.fileName,
      fileType: file.fileType,
      fileSize: file.fileSize
    });
    this.updateModalVisible = true;
  }

  // 处理创建文件
  handleCreateFile() {
    if (this.createForm.invalid) {
      this.message.error('请填写完整信息');
      return;
    }
    
    const input = this.createForm.value;
    this.fileService.create(input).subscribe({
      next: () => {
        this.message.success('文件创建成功');
        this.createModalVisible = false;
        this.getFiles();
      },
      error: (error) => {
        console.error('创建文件失败:', error);
        this.message.error('创建文件失败');
      }
    });
  }

  // 处理更新文件
  handleUpdateFile() {
    if (this.updateForm.invalid || !this.currentFile) {
      this.message.error('请填写完整信息');
      return;
    }
    
    const input = this.updateForm.value;
    this.fileService.update(this.currentFile.id, input).subscribe({
      next: () => {
        this.message.success('文件更新成功');
        this.updateModalVisible = false;
        this.getFiles();
      },
      error: (error) => {
        console.error('更新文件失败:', error);
        this.message.error('更新文件失败');
      }
    });
  }

  // 删除文件
  deleteFile(id: string, fileName: string) {
    this.modal.confirm({
      nzTitle: '确认删除',
      nzContent: `确定要删除文件 "${fileName}" 吗？`,
      nzOkText: '确定',
      nzCancelText: '取消',
      nzOkType: 'primary',
      nzOnOk: () => {
        this.fileService.delete(id).subscribe({
          next: () => {
            this.message.success('文件删除成功');
            this.getFiles();
          },
          error: (error) => {
            console.error('删除文件失败:', error);
            this.message.error('删除文件失败');
          }
        });
      }
    });
  }

  // 格式化文件大小
  formatFileSize(size: number): string {
    if (size === 0) return '0 B';
    const k = 1024;
    const sizes = ['B', 'KB', 'MB', 'GB', 'TB'];
    const i = Math.floor(Math.log(size) / Math.log(k));
    return parseFloat((size / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
  }
}