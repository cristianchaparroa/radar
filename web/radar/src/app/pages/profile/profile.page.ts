import { Component, OnInit } from '@angular/core';


@Component({
  selector: 'app-profile',
  templateUrl: './profile.page.html',
  styleUrls: ['./profile.page.scss'],
})
export class ProfilePage implements OnInit {

  public form = [
    { val: 'Ustede esta enfermo de COVID-19', isChecked: false },
  ];

  constructor() {}

  ngOnInit() {}

}
