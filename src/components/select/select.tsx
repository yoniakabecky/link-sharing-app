import {
  component$,
  useStylesScoped$,
  type PropFunction,
} from "@builder.io/qwik";

import { Icon } from "~/components/icon/icon";
import styles from "./select.css?inline";

type SelectProps = {
  name: string;
  label?: string;
  placeholder?: string;
  value?: string;
  error: string;
  required?: boolean;
  ref: PropFunction<(element: Element) => void>;
  onInput$: PropFunction<(event: Event, element: HTMLSelectElement) => void>;
  onChange$: PropFunction<(event: Event, element: HTMLSelectElement) => void>;
  onBlur$: PropFunction<(event: Event, element: HTMLSelectElement) => void>;
  options: { label: string; value: string }[];
};

export const Select = component$(
  ({ value, options, label, error, ...props }: SelectProps) => {
    const { name, required, placeholder } = props;
    useStylesScoped$(styles);

    return (
      <div>
        {label && (
          <label for={name}>
            {label} {required && <span>*</span>}
          </label>
        )}
        <div class="relative">
          <select
            {...props}
            id={name}
            aria-invalid={!!error}
            aria-errormessage={`${name}-error`}
            class={[value === "" && "empty"]}
          >
            <option value="" disabled hidden selected={!value}>
              {placeholder ?? "Select an Option"}
            </option>
            {options.map(({ label, value: optionValue }) => (
              <option
                key={optionValue}
                value={optionValue}
                selected={optionValue === value}
              >
                {label}
              </option>
            ))}
          </select>
          <Icon name="down" class="arrow" size={20} />
        </div>
        {error && (
          <div id={`${name}-error`} class="error-message">
            {error}
          </div>
        )}
      </div>
    );
  }
);
