import { CommonModule } from "@angular/common";
import { NgModule } from "@angular/core";
import { MaterialModule } from "src/app/shared/modules/material/material.module";
import { CashBookEntryMainCompoent } from "../../components/cash-book-entry-main/cash-book-entry-main.component";

@NgModule({
    declarations: [
        CashBookEntryMainCompoent
    ],
    imports: [
        CommonModule,
        MaterialModule
    ],
    exports: [
        CashBookEntryMainCompoent
    ]
})
export class CashBookEntryModule {}