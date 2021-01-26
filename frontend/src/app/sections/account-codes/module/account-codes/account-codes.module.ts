import { CommonModule } from "@angular/common";
import { NgModule } from "@angular/core";
import { ReactiveFormsModule } from "@angular/forms";
import { MaterialModule } from "src/app/shared/modules/material/material.module";
import { AccountCodesMainComponent } from "../../components/account-codes-main.component";

@NgModule({
    declarations: [
        AccountCodesMainComponent
    ],
    imports: [
        CommonModule,
        MaterialModule,
        ReactiveFormsModule
    ],
    exports: [
        AccountCodesMainComponent
    ]
})
export class AccountCodesModule {}