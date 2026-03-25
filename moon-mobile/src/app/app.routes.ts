import { Routes } from '@angular/router';

export const routes: Routes = [
  {
    path: 'records',
    loadComponent: () => import('./records/records.page').then((m) => m.RecordsPage),
  },
  {
    path: 'home',
    loadComponent: () => import('./home/home.page').then((m) => m.HomePage),
  },
  {
    path: '',
    redirectTo: 'records',
    pathMatch: 'full',
  },
];
