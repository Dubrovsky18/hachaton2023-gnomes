import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AuthTitleComponent } from './components/auth-title/auth-title.component';
import { LoginPageComponent } from './pages/login-page/login-page.component';
import { SelectRolePageComponent } from './pages/select-role-page/select-role-page.component';
import { AuthRoutingModule } from './auth-routing.module';
import { AuthLoginFormComponent } from './components/auth-login-form/auth-login-form.component';



@NgModule({
  declarations: [
    AuthTitleComponent,
    LoginPageComponent,
    SelectRolePageComponent,
    AuthLoginFormComponent,
  ],
  imports: [
    CommonModule,
    AuthRoutingModule
  ]
})
export class AuthModule { }
