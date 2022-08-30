import { Component } from '@angular/core';
import { HttpHeaders, HttpClient, HttpErrorResponse } from '@angular/common/http';
import { catchError, throwError } from 'rxjs';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  constructor(private http: HttpClient){}

  title = 'projectOne';

  ngOnInit() {

  }

  enviarDatos() {
    const headers = new HttpHeaders()
      .append("Authorization", "Bearer " + 'aasdjkqweioasdaslkasd')
      .append("Content-type", "application/json");

    const httpOptions = {
      headers
    };

    const externalUrl = 'http://....';
    const dataToPut = 'Usually, it will be an object, not a string';

    this.http.post<any>(externalUrl, dataToPut, httpOptions)
      .pipe(
        catchError(this.handleError)
      );
  }

  handleError(error: HttpErrorResponse) {
    if (error.error instanceof ErrorEvent) {
      // A client-side or network error occurred. Handle it accordingly.
      console.error('An error occurred:', error.error.message);
    } else {
      // The backend returned an unsuccessful response code.
      // The response body may contain clues as to what went wrong,
      console.error(
        `Backend returned code ${error.status}, ` +
        `body was: ${error.error}`);
    }
    // return an observable with a user-facing error message
    return throwError(
      'Something bad happened; please try again later.');
  };
     

}
