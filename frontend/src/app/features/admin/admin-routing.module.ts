import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CreateUserPageComponent } from './pages/create-user-page/create-user-page.component';

const routes: Routes = [
  { path: 'admin/create-user', component: CreateUserPageComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AdminRoutingModule {}
