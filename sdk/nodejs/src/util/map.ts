export type MapType<V> = { [key: string]: V }

export function filterMap<V>(obj: MapType<V>, filterFunc: (key: string, value: V) => boolean): MapType<V> {
  const result: MapType<V> = {}

  for (const [key, value] of Object.entries(obj)) {
    if (filterFunc(key, value)) {
      result[key] = value
    }
  }

  return result
}
