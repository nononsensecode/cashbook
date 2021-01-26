import { Component } from "@angular/core";
import { HeadingService } from "src/app/shared/services/heading/heading.service";

@Component({
    selector: 'app-settings-main',
    templateUrl: './settings-main.component.html',
    styleUrls: ['./settings-main.component.css']
})
export class SettingsMainComponent { 
    constructor(private headingService: HeadingService) {
        this.headingService.setHeading('Settings');
    }
}