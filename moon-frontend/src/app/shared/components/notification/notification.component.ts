import { FontAwesomeModule, IconDefinition } from '@fortawesome/angular-fontawesome';
import { faExclamationCircle, faX } from '@fortawesome/free-solid-svg-icons';
import { Notification, NotificationService } from '../../service/notification/notification.service';
import { Component, inject, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { merge, Observable, scan, Subject } from 'rxjs';

@Component({
  selector: 'app-notification',
  standalone: true,
  imports: [CommonModule, FontAwesomeModule],
  templateUrl: './notification.component.html',
})
export class NotificationComponent implements OnInit {
  private notificationService = inject(NotificationService);
  private remove$ = new Subject<string>();

  public fax = faX;
  public faExclamationCircle = faExclamationCircle;

  public state$: Observable<{ displaying: Notification[]; waiting: Notification[] }>;

  constructor() {
    this.state$ = merge(this.notificationService.getNotification(), this.remove$).pipe(
      scan(
        (acc, value) => {
          if (typeof value === 'string') {
            // 处理删除通知的逻辑
            const filtered = acc.displaying.filter((n) => n.id !== value);
            const [next, ...restWaiting] = acc.waiting;
            return {
              displaying: next ? [...filtered, next] : filtered,
              waiting: next ? restWaiting : acc.waiting
            };
          } else {
            // 处理添加通知的逻辑
            if (acc.displaying.length < 5) {
              return {
                ...acc,
                displaying: [...acc.displaying, value]
              };
            } else {
              return {
                ...acc,
                waiting: [...acc.waiting, value]
              };
            }
          }
        },
        { displaying: [] as Notification[], waiting: [] as Notification[] }
      )
    );
  }

  public ngOnInit(): void {
    // 组件初始化逻辑
  }

  public removeNotification(id: string): void {
    this.remove$.next(id);
  }
}
