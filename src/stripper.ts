import * as cheerio from 'cheerio';

export default function stripHtml(html: string): string {

    const $ = cheerio.load(html);

    const $mensen  = $('.group').find('.element');
    console.log($mensen.text() + " length: " + $mensen.length);

    const elements = $('.elements:eq(0)');
    console.log(elements.text());
    return $.text();
}

