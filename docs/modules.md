[mensa-api](README.md) / Exports

# mensa-api

## Table of contents

### Functions

- [stripCampus](modules.md#stripcampus)
- [stripMensa](modules.md#stripmensa)

## Functions

### stripCampus

▸ **stripCampus**(`html`): `JSON`

This function strips the html from the campus and Bundesland page and returns a JSON object

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `html` | `string` | The html of the campus or Bundesland page |

#### Returns

`JSON`

JSON object

**`Example`**

```ts
stripCampus(html);
```

#### Defined in

[stripper.ts:141](https://github.com/LeRoid-hub/Mensa-API/blob/5e9a6b1/src/stripper.ts#L141)

___

### stripMensa

▸ **stripMensa**(`html`): `string`

This function strips the html from the mensa page and returns a JSON object

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `html` | `string` | The html of the mensa page |

#### Returns

`string`

JSON object

**`Example`**

```ts
stripMensa(html);
```

#### Defined in

[stripper.ts:168](https://github.com/LeRoid-hub/Mensa-API/blob/5e9a6b1/src/stripper.ts#L168)
