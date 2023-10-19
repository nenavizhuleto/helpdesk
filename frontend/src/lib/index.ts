// place files you want to import through the `$lib` alias in this folder.

export type Property<K> = K extends any ? keyof K : never
