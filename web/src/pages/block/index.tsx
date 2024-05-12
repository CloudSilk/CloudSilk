import '@blocksuite/presets/themes/affine.css';
import { createEmptyPage, DocEditor } from '@blocksuite/presets';

const page = createEmptyPage().init();
const editor = new DocEditor();
editor.page = page;
document.body.appendChild(editor);