import { Injectable } from '@angular/core';
import {
  HttpClient,
  HttpParams,
} from '@angular/common/http';
import { Observable } from 'rxjs';
import { IUser } from 'src/app/core/models/user';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

 constructor(private http: HttpClient) {}

  users: IUser[] = [];

  googleAuth() {
    this.http.get('https://a26980b59d5fb699977bebf5f8e93afb.serveo.net/api/v1/oauth2/login');
  }
}
