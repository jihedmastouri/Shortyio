import { useState } from "react";

const useLocalStorage = <T>(key: string, initialValue?: T) => {
  const [isAvailable, setAvailale] = useState(false);

  const [storedValue, setStoredValue] = useState<T>(() => {
    try {
      const item = window.localStorage.getItem(key);

      if (item) {
        setAvailale(true);
        return JSON.parse(item);
      } else return initialValue;
    } catch (error) {
      return initialValue;
    }
  });

  const setValue = (value: T | Function) => {
    try {
      const valueToStore =
        value instanceof Function ? value(storedValue) : value;

      setAvailale(true);
      setStoredValue(valueToStore);
      window.localStorage.setItem(key, JSON.stringify(valueToStore));
    } catch (error) {
      console.log(error);
    }
  };

  return [storedValue, setValue, isAvailable] as const;
};

export default useLocalStorage;
