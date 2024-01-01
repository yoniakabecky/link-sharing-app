import {
  component$,
  useStylesScoped$,
  type PropFunction,
} from "@builder.io/qwik";

import styles from "./textfield.css?inline";

type TextInputProps = {
  name: string;
  type?: "text" | "email" | "tel" | "password" | "url" | "date";
  label?: string;
  placeholder?: string;
  value?: string;
  error: string;
  required?: boolean;
  ref: PropFunction<(element: Element) => void>;
  onInput$: PropFunction<(event: Event, element: HTMLInputElement) => void>;
  onChange$: PropFunction<(event: Event, element: HTMLInputElement) => void>;
  onBlur$: PropFunction<(event: Event, element: HTMLInputElement) => void>;
};

export const TextField = component$(
  ({ label, error, type, ...props }: TextInputProps) => {
    const { name, required } = props;
    useStylesScoped$(styles);

    return (
      <div>
        {label && (
          <label for={name}>
            {label} {required && <span>*</span>}
          </label>
        )}
        <input
          {...props}
          id={name}
          type={type ?? "text"}
          aria-invalid={!!error}
          aria-errormessage={`${name}-error`}
        />
        {error && (
          <div id={`${name}-error`} class="error-message">
            {error}
          </div>
        )}
      </div>
    );
  }
);
