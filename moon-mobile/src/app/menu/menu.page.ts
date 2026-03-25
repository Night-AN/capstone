import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { IonHeader, IonToolbar, IonTitle, IonContent, IonButton, IonList, IonItem } from '@ionic/angular/standalone';
import { Router } from '@angular/router';

@Component({
  selector: 'app-menu',
  templateUrl: './menu.page.html',
  styleUrls: ['./menu.page.scss'],
  imports: [
    CommonModule,
    IonHeader, IonToolbar, IonTitle, IonContent,
    IonButton, IonList, IonItem
  ],
})
export class MenuPage {

  constructor(private router: Router) {}

  goToProcurement(): void {
    this.router.navigate(['/procurement']);
  }

  goToAcceptance(): void {
    this.router.navigate(['/acceptance']);
  }
}
