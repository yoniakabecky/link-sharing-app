import type { IconName } from "~/components/icon/icon";

export type PlatformType =
  | "custom"
  | "discord"
  | "dribbble"
  | "facebook"
  | "github"
  | "instagram"
  | "linkedin"
  | "portfolio"
  | "social_x"
  | "tiktok"
  | "youtube";

export type Platform = {
  type: PlatformType;
  label: string;
  url: string;
  color: string;
  icon: IconName;
};

export const platforms: Record<PlatformType, Platform> = {
  discord: {
    type: "discord",
    label: "Discord",
    url: "https://discord.com/",
    color: "#5865F2",
    icon: "discord",
  },
  dribbble: {
    type: "dribbble",
    label: "Dribbble",
    url: "https://dribbble.com/",
    color: "#EA4C89",
    icon: "dribbble",
  },
  facebook: {
    type: "facebook",
    label: "Facebook",
    url: "https://www.facebook.com/",
    color: "#3B5998",
    icon: "facebook",
  },
  github: {
    type: "github",
    label: "GitHub",
    url: "https://github.com/",
    color: "#333333",
    icon: "github",
  },
  instagram: {
    type: "instagram",
    label: "Instagram",
    url: "https://www.instagram.com/",
    color: "#E4405F",
    icon: "instagram",
  },
  linkedin: {
    type: "linkedin",
    label: "LinkedIn",
    url: "https://www.linkedin.com/",
    color: "#0077B5",
    icon: "linkedin",
  },
  portfolio: {
    type: "portfolio",
    label: "Portfolio",
    url: "https://www.example.com/",
    color: "#000000",
    icon: "web",
  },
  social_x: {
    type: "social_x",
    label: "X (Twitter)",
    url: "https://twitter.com/",
    color: "#000000",
    icon: "social_x",
  },
  tiktok: {
    type: "tiktok",
    label: "TikTok",
    url: "https://www.tiktok.com/",
    color: "#010101",
    icon: "tiktok",
  },
  youtube: {
    type: "youtube",
    label: "YouTube",
    url: "https://www.youtube.com/",
    color: "#FF0000",
    icon: "youtube",
  },
  custom: {
    type: "custom",
    label: "",
    url: "",
    color: "#000000",
    icon: "web",
  },
};
