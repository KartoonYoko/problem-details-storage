import {Component, OnInit} from '@angular/core';
import {MatCardModule} from '@angular/material/card';
import {NgForOf, NgIf} from '@angular/common';
import {MatButton} from '@angular/material/button';
import {debounceTime, distinctUntilChanged, of, Subject, switchMap, tap} from 'rxjs';
import {MatProgressBar} from '@angular/material/progress-bar';

interface bucket {
    name: string
    description: string
}

@Component({
    selector: 'app-buckets-list',
    imports: [
        MatCardModule,
        NgForOf,
        MatButton,
        MatProgressBar,
        NgIf
    ],
    templateUrl: './buckets-list.component.html',
    styleUrl: './buckets-list.component.css'
})
export class BucketsListComponent implements OnInit {
    private searchText = new Subject<string>();

    private bucketsSample: bucket[] = [
        {
            name: 'Приложение доставки',
            description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.'
        },
        {
            name: 'Сервис заказов',
            description: 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.'
        },
        {
            name: 'Сервис уведомлений',
            description: 'It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout.'
        }
    ];

    isLoading = false;

    filteredBucketsList: bucket[] = this.bucketsSample;

    ngOnInit() {
        this.searchText.pipe(
            tap(() => this.isLoading = true),
            debounceTime(500),
            distinctUntilChanged(),
            switchMap(searchText => of(null).pipe(
                tap(() => {
                    const lowSearchText = searchText.toLowerCase();

                    this.filteredBucketsList = this.bucketsSample
                        .filter(x => x.name.toLowerCase().includes(lowSearchText));

                    this.isLoading = false;
                })
            ))
        ).subscribe();
    }

    search(searchText: string) {
        this.searchText.next(searchText);
    }
}
