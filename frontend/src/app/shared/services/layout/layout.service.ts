import { BreakpointObserver, Breakpoints } from "@angular/cdk/layout";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { map } from 'rxjs/operators';

@Injectable({
    providedIn: 'root'
})
export class LayoutService {
    constructor(private observer: BreakpointObserver) {}

    public isMobile(): Observable<boolean> {
        return this.observer.observe([
            Breakpoints.Handset,
            Breakpoints.TabletPortrait
        ]).pipe(
            map(results => {
                return results.matches;
            })
        );
    }
}