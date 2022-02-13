import {Component, OnInit} from '@angular/core';
import {MatTableDataSource} from "@angular/material/table";
import {Question} from "./question";
import {ApiService} from "./api.service";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent implements OnInit {
  title = 'AskMeAnythingUI';
  question: Question[] = [];
  // columns we will show on the table
  public displayedColumns = ['id', 'question', 'answer'];
  // the source where we will get the data
  public dataSource = new MatTableDataSource<Question>();

  // dependency injection
  constructor(private questionApiService: ApiService) {
  }

  ngOnInit() {
    // call this method on component load
    this.getQuestionsInformation();
  }

  /**
   * This method returns questions details
   */
  getQuestionsInformation() {
    this.questionApiService.getQuestionsInformation()
      .subscribe((res)=>{
        console.log(res);
        this.dataSource.data = res.data.data;
      })
  }
}
