import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CheckboxComponent } from './components/checkbox/checkbox.component';
import { ProfileComponent } from './components/profile/profile.component';
import { HttpClientModule } from '@angular/common/http';
import { OAuthModule } from 'angular-oauth2-oidc';



@NgModule({
  declarations: [
    CheckboxComponent,
    ProfileComponent,
  ],
  imports: [
    CommonModule,
    HttpClientModule,
    OAuthModule.forRoot()
  ],
  exports: [
    CheckboxComponent,
    ProfileComponent,
  ],
})
export class SharedModule { }
