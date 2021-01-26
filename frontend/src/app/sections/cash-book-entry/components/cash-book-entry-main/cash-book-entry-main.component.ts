import { Component } from "@angular/core";
import { HeadingService } from "src/app/shared/services/heading/heading.service";

@Component({
    selector: 'app-cash-book-entry-main',
    templateUrl: './cash-book-entry-main.component.html',
    styleUrls: ['./cash-book-entry-main.component.css']
})
export class CashBookEntryMainCompoent {
    constructor(private headingService: HeadingService) {
        this.headingService.setHeading('Cashbook Entry');
    }
}