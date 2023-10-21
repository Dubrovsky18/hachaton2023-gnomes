import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { SelectRolePageComponent } from './pages/select-role-page/select-role-page.component';
import { LoginPageComponent } from './pages/login-page/login-page.component';

const routes: Routes = [
  { path: '', component: SelectRolePageComponent },
  { path: 'login', component: LoginPageComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AuthRoutingModule {}
