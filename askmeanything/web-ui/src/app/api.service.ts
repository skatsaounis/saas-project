import { Injectable } from '@angular/core';
// import question interface
import {Question} from "./question";
import {QResponse} from "./question";
// import this to make http requests
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
// we've defined our base url here in the env
import {environment} from "./../environments/environment";

@Injectable({
  providedIn: 'root'
})

export class ApiService {

  constructor(private httpClient: HttpClient) { }

  /**
   * This method returns questions details
   */
  getQuestionsInformation(): Observable<QResponse>{
    return this.httpClient.get<QResponse>(`${environment.questionsURL}`);
  }
}
