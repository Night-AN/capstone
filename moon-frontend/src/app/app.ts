import { Component, inject, signal } from '@angular/core';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { faCoffee } from '@fortawesome/free-solid-svg-icons';

import { RouterOutlet } from '@angular/router';
import { NotificationService } from './shared/service/notification/notification.service';
import { NotificationComponent } from './shared/components/notification/notification.component';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet,
    FontAwesomeModule,
    NotificationComponent
  ],
  templateUrl: './app.html',
})
export class App {
  protected readonly title = signal('moon-frontend');
  public notify = inject(NotificationService)
    faCoffee = faCoffee;

  
}
