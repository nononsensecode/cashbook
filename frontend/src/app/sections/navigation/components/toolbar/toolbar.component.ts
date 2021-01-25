import { EventEmitter, OnInit, Output } from "@angular/core";
import { Observable } from "rxjs";
import { LayoutService } from "src/app/shared/services/layout/layout.service";

export class ToolbarComponent implements OnInit {

    public isMobile$: Observable<boolean>;

    @Output() mainMenuClick = new EventEmitter<void>();
    
    constructor(private layoutService: LayoutService) {}

    ngOnInit() {
        this.isMobile$ = this.layoutService.isMobile();
    }

    onMainMenuClick() {
        this.mainMenuClick.emit();
    }
}