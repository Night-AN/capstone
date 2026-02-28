import { CommonModule } from '@angular/common';
import { Component, EventEmitter, Input, Output, signal } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';

@Component({
  selector: 'app-organization-tree-node',
  standalone: true,
  imports: [
    CommonModule,
    MatButtonModule,
    MatIconModule
  ],
  templateUrl: './organization-tree-node.component.html',
  styleUrl: './organization-tree-node.component.scss'
})
export class OrganizationTreeNodeComponent {
  @Input() organization: any;
  @Input() loading: boolean = false;
  @Output() view = new EventEmitter<string>();
  @Output() edit = new EventEmitter<string>();
  @Output() delete = new EventEmitter<string>();

  isExpanded = signal<boolean>(false);

  toggleExpand(): void {
    if (this.organization.children && this.organization.children.length > 0) {
      this.isExpanded.set(!this.isExpanded());
    }
  }

  onView(): void {
    this.view.emit(this.organization.organization_id);
  }

  onEdit(): void {
    this.edit.emit(this.organization.organization_id);
  }

  onDelete(): void {
    this.delete.emit(this.organization.organization_id);
  }
}
