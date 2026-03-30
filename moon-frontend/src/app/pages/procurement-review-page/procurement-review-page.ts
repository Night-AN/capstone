import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzCardModule } from 'ng-zorro-antd/card';
import { NzFormModule } from 'ng-zorro-antd/form';
import { NzInputModule } from 'ng-zorro-antd/input';
import { NzTableModule } from 'ng-zorro-antd/table';
import { NzModalModule } from 'ng-zorro-antd/modal';
import { NzMessageService } from 'ng-zorro-antd/message';
import { CommonModule } from '@angular/common';
import { ProcurementReviewService } from 'src/app/services/procurement-review.service';

interface ProcurementReview {
  id: string;
  review_name: string;
  review_code: string;
  review_description: string;
  review_flag: string;
  created_at: string;
  updated_at: string | null;
}

enum ModalTitle {
  Create = '创建评审记录',
  Edit = '编辑评审记录'
}

@Component({
  selector: 'app-procurement-review-page',
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    NzButtonModule,
    NzCardModule,
    NzFormModule,
    NzInputModule,
    NzTableModule,
    NzModalModule
  ],
  templateUrl: './procurement-review-page.html',
  styleUrl: './procurement-review-page.less',
  providers: [NzMessageService]
})
export class ProcurementReviewPage implements OnInit {
  reviewForm: FormGroup;
  reviews: ProcurementReview[] = [];
  originalReviews: ProcurementReview[] = [];

  isEditing = false;
  currentReviewId: string | null = null;

  visible = false;
  modalTitle = ModalTitle.Create;
  isLoading = false;
  searchKeyword: string = '';

  constructor(
    private fb: FormBuilder,
    private message: NzMessageService,
    private reviewService: ProcurementReviewService
  ) {
    this.reviewForm = this.fb.group({
      review_name: ['', [Validators.required]],
      review_code: ['', [Validators.required]],
      review_description: ['', [Validators.required]],
      review_flag: ['', [Validators.required]]
    });
  }

  ngOnInit(): void {
    this.loadReviews();
  }

  loadReviews(): void {
    this.reviewService.getList().subscribe(reviews => {
      this.reviews = reviews[0].edges.map((edge: any) => ({
        id: edge.node.id,
        review_name: edge.node.review_name,
        review_code: edge.node.review_code,
        review_description: edge.node.review_description,
        review_flag: edge.node.review_flag,
        created_at: edge.node.created_at,
        updated_at: edge.node.updated_at
      }));
      this.originalReviews = [...this.reviews];
    });
  }

  searchReviews(): void {
    if (this.searchKeyword) {
      const keyword = this.searchKeyword.toLowerCase();
      this.reviews = this.originalReviews.filter(review => 
        review.review_name.toLowerCase().includes(keyword) ||
        review.review_code.toLowerCase().includes(keyword) ||
        review.review_description.toLowerCase().includes(keyword) ||
        review.review_flag.toLowerCase().includes(keyword)
      );
    } else {
      this.reviews = [...this.originalReviews];
    }
  }

  createReview(): void {
    this.resetForm();
    this.modalTitle = ModalTitle.Create;
    this.visible = true;
  }

  editReview(review: ProcurementReview): void {
    this.isEditing = true;
    this.currentReviewId = review.id;
    this.modalTitle = ModalTitle.Edit;
    this.reviewForm.patchValue({
      review_name: review.review_name,
      review_code: review.review_code,
      review_description: review.review_description,
      review_flag: review.review_flag
    });
    this.visible = true;
  }

  deleteReview(id: string): void {
    this.reviewService.delete(id).subscribe(() => {
      this.loadReviews();
      this.message.success('评审记录删除成功');
    });
  }

  handleCancel(): void {
    this.visible = false;
    this.resetForm();
  }

  handleOk(): void {
    if (this.reviewForm.valid) {
      this.isLoading = true;
      setTimeout(() => {
        if (this.isEditing && this.currentReviewId) {
          this.updateReview();
        } else {
          this.createReviewSubmit();
        }
        this.isLoading = false;
        this.visible = false;
      }, 500);
    } else {
      Object.values(this.reviewForm.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }

  createReviewSubmit(): void {
    this.reviewService.create(this.reviewForm.value).subscribe(() => {
      this.loadReviews();
      this.message.success('评审记录创建成功');
      this.resetForm();
    });
  }

  updateReview(): void {
    if (this.currentReviewId) {
      this.reviewService.update(this.currentReviewId, this.reviewForm.value).subscribe(() => {
        this.loadReviews();
        this.message.success('评审记录更新成功');
        this.resetForm();
      });
    }
  }

  resetForm(): void {
    this.reviewForm.reset();
    this.isEditing = false;
    this.currentReviewId = null;
  }
}