import { ObserversModule } from "@angular/cdk/observers";
import { PlatformModule } from "@angular/cdk/platform";
import { CommonModule } from "@angular/common";
import { NgModule } from "@angular/core";
import { FlexLayoutModule } from "@angular/flex-layout";
import { MatButtonModule, MatCardModule, MatCheckboxModule, MatDividerModule, MatIconModule,
     MatInputModule, MatListModule, MatSelectModule, MatSidenavModule,
     MatSnackBarModule, MatTableModule, MatToolbarModule, MatTooltipModule } from "@angular/material";

@NgModule({
    imports: [
        CommonModule,
        MatInputModule,
        MatButtonModule,
        MatCardModule,
        MatToolbarModule,
        MatSidenavModule,
        MatIconModule,
        MatListModule,
        MatSnackBarModule,
        MatDividerModule,
        MatSelectModule,
        MatTableModule,
        MatTooltipModule,
        MatCheckboxModule,
        FlexLayoutModule,
        PlatformModule,
        ObserversModule
    ],
    exports: [
        CommonModule,
        MatInputModule,
        MatButtonModule,
        MatCardModule,
        MatToolbarModule,
        MatSidenavModule,
        MatIconModule,
        MatListModule,
        MatSnackBarModule,
        MatDividerModule,
        MatSelectModule,
        MatTableModule,
        MatTooltipModule,
        MatCheckboxModule,
        FlexLayoutModule,
        PlatformModule,
        ObserversModule
    ]
})
export class MaterialModule{}