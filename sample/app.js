const API_BASE_URL = 'http://localhost:8080/api';

// Content Types と Fields を取得する関数
async function fetchContentTypesAndFields() {
  try {
    // Content Types の取得
    const contentTypesResponse = await fetch(`${API_BASE_URL}/content-types`);
    const contentTypes = await contentTypesResponse.json();

    // Fields の取得
    const fieldsResponse = await fetch(`${API_BASE_URL}/fields`);
    const fields = await fieldsResponse.json();

    // Content Type ごとに Fields をグループ化
    const contentTypesWithFields = contentTypes.map(contentType => ({
      ...contentType,
      fields: fields.filter(field => field.contentTypeID === contentType.id)
    }));

    displayContentTypes(contentTypesWithFields);
  } catch (error) {
    displayError('データの取得中にエラーが発生しました: ' + error.message);
  }
}

// Content Types と Fields を画面に表示する関数
function displayContentTypes(contentTypesWithFields) {
  const container = document.getElementById('content-types');
  container.innerHTML = ''; // ローディング表示をクリア

  if (contentTypesWithFields.length === 0) {
    container.innerHTML = '<div class="error">コンテンツタイプが見つかりません。</div>';
    return;
  }

  contentTypesWithFields.forEach(contentType => {
    const contentTypeElement = document.createElement('div');
    contentTypeElement.className = 'content-type';

    contentTypeElement.innerHTML = `
            <h2>${contentType.displayName || contentType.name}</h2>
            ${contentType.description ? `<p>${contentType.description}</p>` : ''}
            <div class="field-list">
                <h3>Fields:</h3>
                ${contentType.fields.length > 0
        ? contentType.fields
          .sort((a, b) => a.order - b.order)
          .map(field => `
                            <div class="field">
                                <strong>${field.label || field.name}</strong>
                                <br>
                                Type: ${field.type}
                                ${field.isRequired ? ' <span style="color: red;">*必須</span>' : ''}
                            </div>
                        `).join('')
        : '<div class="error">フィールドが定義されていません。</div>'
      }
            </div>
        `;

    container.appendChild(contentTypeElement);
  });
}

// エラーメッセージを表示する関数
function displayError(message) {
  const container = document.getElementById('content-types');
  container.innerHTML = `<div class="error">${message}</div>`;
}

// ページ読み込み時にデータを取得
document.addEventListener('DOMContentLoaded', fetchContentTypesAndFields); 
