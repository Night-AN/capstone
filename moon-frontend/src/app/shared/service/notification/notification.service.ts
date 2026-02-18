import { Injectable, OnInit } from '@angular/core';
import { Observable, Subject } from 'rxjs';

export interface Notification {
  id?: string;
  title?: string;
  message: string;
  type: 'success' | 'error' | 'warning' | 'info';
  duration?: number;
}

@Injectable({
  providedIn: 'root',
})
export class NotificationService {
  
  private notificationSubject = new Subject<Notification>();

  public getNotification(): Observable<Notification> {
    return this.notificationSubject.asObservable();
  }

  public show(notification: Notification): void {
    notification.id = `n-${Date.now()}-${Math.random().toString(36).substring(2, 9)}`;
    this.notificationSubject.next({ duration: 3000, ...notification });
  }

  success(message: string, title?: string, duration?: number): void {
    this.show({ title, message, type: 'success', duration });
  }

  error(message: string, title?: string, duration?: number): void {
    this.show({ title, message, type: 'error', duration });
  }

  warning(message: string, title?: string, duration?: number): void {
    this.show({ title, message, type: 'warning', duration });
  }

  info(message: string, title?: string, duration?: number): void {
    this.show({ title, message, type: 'info', duration });
  }
}
