import * as cheerio from 'cheerio';

class Campus {
    name: string;
    mensen: Mensa[];

    constructor(name: string) {
        this.name = name;
        this.mensen = [];
    }
    addMensa(mensa: Mensa) {
        this.mensen.push(mensa);
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
    return "Mensa";
}

