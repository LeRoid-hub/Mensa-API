import * as cheerio from 'cheerio';

class Campus {
    name: string;
    locations: Mensa[];

    constructor(name: string) {
        this.name = name;
        this.locations = [];
    }
    addMensa(mensa: Mensa) {
        this.locations.push(mensa);
    }
    toString() {
        let str = this.name + "\n";
        this.locations.forEach((elem) => {
            str += "  -> " + elem.name + " | " + elem.url + "\n";
        });
        return str;
    }
}

class Mensa {
    name: string;
    url: string;

    constructor(name: string, url: string) {
        this.name = name;
        this.url = url;
    }
}

class Menu {
    name: string;
    meals: Meal[];

    constructor(name: string) {
        this.name = name;
        this.meals = [];
    }
    addMeal(meal: Meal) {
        this.meals.push(meal);
    }

    toString() {
        let str = this.name + "\n";
        this.meals.forEach((elem) => {
            str += "  -> " + elem.name + " | " + elem.price + " | " + elem.atributs + "\n";
        });
        return str;
    }
}

class Meal {
    name: string;
    price: string;
    atributs: string[];

    constructor(name: string, price: string,  atributs: string[] ) {
        this.name = name;
        this.price = price;
        this.atributs = atributs;
    }
}

export function stripCampus(html: string): JSON {

    const $ = cheerio.load(html);

    let ort :Campus[] = [];
    const $campus = $('.tagged').find('.group');

    $campus.each((i, elem) => {
        let campus = new Campus($(elem).find('h2').text());
        //console.log($(elem).find('h2').text() + "\n");
        $(elem).find('a').each((i, elem) => {
            const link = $(elem).attr('href')?.replace("/index.html","");
            campus.addMensa(new Mensa($(elem).text(), link ?? ""));
            //console.log("  -> " + $(elem).text() + " | " + $(elem).attr('href')?.replace("/index.html","")  + "\n");
        });
        ort.push(campus);
    });
    return JSON.parse(JSON.stringify(ort));
}

export function stripMensa(html: string): string {
    const $ = cheerio.load(html);

    let menus :Menu[] = [];

    const $menu = $('.aw-meal-category'); 
    $menu.each((i, elem) => {
        const menu = new Menu($(elem).find('h3').text());

        //meal
        $(elem).find('.aw-meal').each((i, elem) => {
            const name = $(elem).find('.aw-meal-description').text();
            const price = $(elem).find('.aw-meal-price').text();
            
            //delete  last time served
            $(elem).find('.aw-meal-attributes').find('.hidden-md').remove();
            
            const atributs :string[] = $(elem).find('.aw-meal-attributes').text().split(" ");

            atributs.forEach((elem, i) => {
                elem = elem.trim();
                atributs[i] = elem;
            });


            menu.addMeal(new Meal(name, price, atributs));
        });
        menus.push(menu);
    });
    if (menus.length === 0) {
        return JSON.parse("No usable Data found");
    }
    return JSON.parse(JSON.stringify(menus));
}

