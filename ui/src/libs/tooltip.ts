export default function Tooltip(node: any, params: any = {}) {
    // @ts-ignore
    if (!window.tippy) return;

    const dataTooltipTemplate = node.getAttribute("data-tooltip-template");

    // @ts-ignore
    const tip = tippy(node, {
        content(reference) {
            const contentText = document.getElementById(dataTooltipTemplate);
            return contentText.innerHTML;
        },
        onShow(instance) {
            const contentText = document.getElementById(dataTooltipTemplate);

            instance.setContent(contentText.innerHTML);
        },
        allowHTML: true,
        arrow: false,
        ...params,
    });

    return {
        update: (newParams) => tip.setProps(...params),
        destroy: () => tip.destroy(),
    };
}