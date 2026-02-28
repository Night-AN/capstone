import { CommonModule } from '@angular/common';
import { Component, EventEmitter, Input, Output, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';

interface TreeSelectItem {
  id: string;
  label: string;
  checked: boolean;
  disabled?: boolean;
  expanded?: boolean;
  children?: TreeSelectItem[];
}

@Component({
  selector: 'app-tree-select',
  imports: [
    CommonModule,
    FormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatIconModule,
    MatCheckboxModule
  ],
  templateUrl: './tree-select.component.html',
  styleUrl: './tree-select.component.scss'
})
export class TreeSelectComponent {
  @Input() label: string = '选择';
  @Input() placeholder: string = '请选择';
  @Input() items: TreeSelectItem[] = [];
  @Output() selectionChange = new EventEmitter<TreeSelectItem[]>();

  isOpen = signal<boolean>(false);
  searchTerm = '';

  get displayValue(): string {
    const selectedItems = this.getAllSelectedItems(this.items);
    if (selectedItems.length === 0) {
      return '';
    }
    if (selectedItems.length > 3) {
      return `${selectedItems.length} 项已选择`;
    }
    return selectedItems.map(item => item.label).join(', ');
  }

  get filteredItems(): TreeSelectItem[] {
    if (!this.searchTerm) {
      return this.items;
    }
    return this.filterItems(this.items, this.searchTerm);
  }

  toggleDropdown(): void {
    this.isOpen.set(!this.isOpen());
  }

  toggleItem(item: TreeSelectItem): void {
    if (item.children && item.children.length > 0) {
      item.expanded = !item.expanded;
    }
  }

  onItemCheck(item: TreeSelectItem): void {
    if (item.children && item.children.length > 0) {
      this.toggleChildren(item, item.checked);
    }
    this.selectionChange.emit(this.getAllSelectedItems(this.items));
  }

  private toggleChildren(item: TreeSelectItem, checked: boolean): void {
    if (item.children) {
      item.children.forEach(child => {
        if (!child.disabled) {
          child.checked = checked;
          if (child.children && child.children.length > 0) {
            this.toggleChildren(child, checked);
          }
        }
      });
    }
  }

  private getAllSelectedItems(items: TreeSelectItem[]): TreeSelectItem[] {
    let selectedItems: TreeSelectItem[] = [];
    items.forEach(item => {
      if (item.checked) {
        selectedItems.push(item);
      }
      if (item.children && item.children.length > 0) {
        selectedItems = [...selectedItems, ...this.getAllSelectedItems(item.children)];
      }
    });
    return selectedItems;
  }

  private filterItems(items: TreeSelectItem[], searchTerm: string): TreeSelectItem[] {
    const filtered: TreeSelectItem[] = [];
    items.forEach(item => {
      if (item.label.toLowerCase().includes(searchTerm.toLowerCase())) {
        filtered.push({ ...item });
      }
      if (item.children && item.children.length > 0) {
        const filteredChildren = this.filterItems(item.children, searchTerm);
        if (filteredChildren.length > 0) {
          filtered.push({
            ...item,
            children: filteredChildren,
            expanded: true
          });
        }
      }
    });
    return filtered;
  }
}
