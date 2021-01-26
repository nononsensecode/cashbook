import { Component } from "@angular/core";
import { HeadingService } from "src/app/shared/services/heading/heading.service";

@Component({
    selector: 'app-dashboard',
    templateUrl: './dashboard-main.component.html',
    styleUrls: ['./dashboard-main.component.css']
})
export class DashboardMainComponent {
    constructor(private headingService: HeadingService) {
        this.headingService.setHeading('Dashboard');
    }
}