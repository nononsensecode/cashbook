import { Component } from "@angular/core";
import { Observable } from "rxjs";
import { HeadingService } from "src/app/shared/services/heading/heading.service";
import { LayoutService } from "src/app/shared/services/layout/layout.service";

@Component({
    selector: 'app-header',
    templateUrl: './header.component.html',
    styleUrls: ['./header.component.css']
})
export class HeaderComponent {
    isMobile$: Observable<boolean>;
    heading$: Observable<string>;

    constructor(
        private layoutService: LayoutService,
        private headingService: HeadingService
    ) {
        this.isMobile$ = this.layoutService.isMobile();
        this.heading$ = this.headingService.heading$;
    }
}