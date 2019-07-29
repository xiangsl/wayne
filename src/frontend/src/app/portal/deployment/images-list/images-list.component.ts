import { AfterViewInit, Component, Inject, OnDestroy, OnInit, ViewChild, Input, EventEmitter, Output } from '@angular/core';
import 'rxjs/add/operator/debounceTime';
import 'rxjs/add/operator/distinctUntilChanged';
import { DOCUMENT, Location } from '@angular/common';
import { FormBuilder, NgForm } from '@angular/forms';
import { EventManager } from '@angular/platform-browser';
import { MessageHandlerService } from '../../../shared/message-handler/message-handler.service';
import { ImagesService } from 'app/shared/client/v1/images.service';


@Component({
  selector: 'wayne-images-list',
  templateUrl: './images-list.component.html',
  styleUrls: ['./images-list.component.scss']
})
export class ImagesListComponent implements OnInit {

  constructor(private imagesService: ImagesService) { }
  public ifShow: any = true;
  apps: any = [];
  page: any = {};
  version_name = '';
  @Output() selectImg = new EventEmitter<string>();
  pageSizeChange(pageSize: number) {
    this.page.currentPage = 1;
    this.page.pageSize = pageSize;
    this.listImages();
  }
  curPageChange(curPage: number) {
    this.page.currentPage = curPage;
    this.listImages();
  }
  listImages() {
    this.page.versionName = this.version_name;
    this.imagesService.listPage(this.page).subscribe(
      response => {
      const data = response.data;
      this.page.totalPage = data.totalPage;
      this.page.totalCount = data.totalCount;
      this.apps = data.list;
     })
  }
  selectImage(app: any) {
    this.selectImg.emit(app.repository_harbor + ':' + app.tag);
  }
  imageDetail(app: any) {
  }
  ngOnInit() {
    this.page.totalPage = 0;
    this.page.totalCount = 0;
    this.page.currentPage = 1;
    this.page.pageSize = 10;
  this.listImages();
  }
}
