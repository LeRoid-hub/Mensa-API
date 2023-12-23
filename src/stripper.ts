import * as cheerio from 'cheerio';

/**
    * This class represents a Campus or Bundesland
    *@class campus
    *@property {string} name The name of the Campus or Bundesland
    *@property {mensa[]} locations The locations of the Campus or Bundesland
*/
class Campus {
    name: string;
    locations: Mensa[];

    /**
        * constructor for campus
        * @param name The name of the Campus or Bundesland
        * @returns Campus
        * @example
        * let campus = new Campus("Campus");
        */
    constructor(name: string) {
        this.name = name;
        this.locations = [];
    }

    /**
        * This function adds a location to the Campus or Bundesland
        * @param mensa The location or to add
        * @returns void
        * @example
        * campus.addMensa(mensa);
        */
    addMensa(mensa: Mensa) {
        this.locations.push(mensa);
    }

    /**
        * This function returns a string representation of the Campus or Bundesland
    */ 
    toString() {
        let str = this.name + "\n";
        this.locations.forEach((elem) => {
            str += "  -> " + elem.name + " | " + elem.url + "\n";
        });
        return str;
    }
}

/**
    * This class represents a mensa
    *@class Mensa
    *@property {string} name The name of the Mensa
    *@property {string} url The url of the mensa
*/
class Mensa {
    name: string;
    url: string;

    /**
        * constructor for mensa
        * @param name The name of the mensa 
        * @param url The url of the Mensa
        * @example
        * let mensa = new Mensa("Mensa", url;
        */
    constructor(name: string, url: string) {
        this.name = name;
        this.url = url;
    }
}

/**
    * This class represents a menus
    *@class menus
    *@property {string} name The name of the menus
    *@property {meal[]} meals The meals of the menus
*/
class Menu {
    name: string;
    meals: Meal[];

    /**
        * constructor for menus
        * @param name The name of the menus
    */
    constructor(name: string) {
        this.name = name;
        this.meals = [];
    }

    /**
        * This function adds a meal to the menus
        * @param meal The meal to add 
        * @returns void
        * @example
        * menu.addMeal(meal);
        */
    addMeal(meal: Meal) {
        this.meals.push(meal);
    }

    /**
        * This function returns a string representation of the menus
        * @returns string 
    */
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

    /**
        * constructor for meals
        * @param name The name of the meal 
        * @param price The price of the meals
        * @param atributs The atributs of the meals
        * @example
        * let meal = new Meal("Pommes", "2,50€", ["vegan", "vegetarisch"]);
        */
    constructor(name: string, price: string,  atributs: string[] ) {
        this.name = name;
        this.price = price;
        this.atributs = atributs;
    }
}

/**
    * This function strips the html from the campus and Bundesland page and returns a JSON object
    * @param html The html of the campus or Bundesland page
    * @returns JSON object
    * @example
    * stripCampus(html);
*/
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

/**
    * This function strips the html from the mensa page and returns a JSON object
    * @param html The html of the mensa page
    * @returns JSON object
    * @example
    * stripMensa(html);
*/
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

