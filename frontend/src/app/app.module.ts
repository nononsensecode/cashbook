import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { APP_BASE_HREF } from '@angular/common';
import { MaterialModule } from './shared/modules/material/material.module';
import { NavigationModule } from './sections/navigation/module/navigation/navigation.module';
import { HeaderFooterModule } from './sections/header-footer/module/header-footer/header-footer.module';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    MaterialModule,
    NavigationModule,
    HeaderFooterModule,
    BrowserAnimationsModule
  ],
  providers: [{provide: APP_BASE_HREF, useValue : '/' }],
  bootstrap: [AppComponent]
})
export class AppModule { }
