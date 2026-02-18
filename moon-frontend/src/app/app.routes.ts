import { Routes } from '@angular/router';
import { AuthModule } from './features/auth/auth.module';
import { UserModule } from './features/user/user.module';
import { OrganizationModule } from './features/organization/organization.module';
import { PermissionModule } from './features/permission/permission.module';
import { RoleModule } from './features/role/role.module';
import { ResourceModule } from './features/resource/resource.module';


export const routes: Routes = [
{
    path:'auth',
    loadChildren:()=>AuthModule
},
{
    path:'',
    loadComponent: () => import('./index.component').then(m => m.IndexComponent),
    children: [
        {
            path:'test',
            loadComponent: () => import('./features/test/test.component').then(m => m.TestComponent)
        },
        {
            path:'dashboard',
            loadComponent: () => import('./features/dashboard/dashboard.component').then(m => m.DashboardComponent)
        },
        {
            path:'users',
            loadChildren:()=>UserModule
        },
        {
            path:'organizations',
            loadChildren:()=>OrganizationModule
        },
        {
            path:'roles',
            loadChildren:()=>RoleModule
        },
        {
            path:'permissions',
            loadChildren:()=>PermissionModule
        },
        {
            path:'resources',
            loadChildren:()=>ResourceModule
        },
        {
            path:'',
            redirectTo:'dashboard',
            pathMatch:'full'
        }
    ]
}
];

