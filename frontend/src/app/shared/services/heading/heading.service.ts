import { Injectable } from "@angular/core";
import { Subject } from "rxjs";

@Injectable({
    providedIn: 'root'
})
export class HeadingService {
    private headingSource = new Subject<string>();
    public heading$ = this.headingSource.asObservable();

    public setHeading(heading: string) {
        this.headingSource.next(heading);
    }
}