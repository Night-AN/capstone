import { ApplicationConfig, provideBrowserGlobalErrorListeners, inject } from '@angular/core';
import { provideRouter } from '@angular/router';

import { routes } from './app.routes';
import { provideClientHydration, withEventReplay, withHttpTransferCacheOptions } from '@angular/platform-browser';
import { icons } from './icons-provider';
import { provideNzIcons } from 'ng-zorro-antd/icon';
import { zh_CN, provideNzI18n } from 'ng-zorro-antd/i18n';
import { HttpHeaders, provideHttpClient, withFetch, withInterceptors } from '@angular/common/http';
import { provideApollo } from 'apollo-angular';
import { HttpLink } from 'apollo-angular/http';
import { InMemoryCache } from '@apollo/client';
import { AuthInterceptor } from '@services/auth.interceptor';
import { SetContextLink } from '@apollo/client/link/context';
import { provideNzConfig } from 'ng-zorro-antd/core/config';

export const appConfig: ApplicationConfig = {
  providers: [
    provideBrowserGlobalErrorListeners(),
    provideRouter(routes),
    provideClientHydration(withEventReplay(), withHttpTransferCacheOptions({
      includePostRequests: true
    })),
    provideNzIcons(icons), provideNzI18n(zh_CN),
    provideHttpClient(
      withInterceptors([AuthInterceptor]), withFetch()
    ), provideApollo(() => {
      const httpLink = inject(HttpLink);
      const token = typeof window !== 'undefined' ? localStorage.getItem('token') : null;

      const authLink = new SetContextLink((preContext,operation)=>{
        return {
          headers: new HttpHeaders({
            "Authorization": token ? `Bearer ${token}` : '',
          })
        };
      });

      return {
        link: authLink.concat(httpLink.create({
          uri: 'http://localhost:8080/query'
        })),
        cache: new InMemoryCache(),
      };
    })
  ]
};
