new Vue({
    // 「el」プロパティーで、Vueの表示を反映する場所＝HTML要素のセレクター（id）を定義
    el: '#appChatMember',

    // data オブジェクトのプロパティの値を変更すると、ビューが反応し、新しい値に一致するように更新
    data: {
        // ログインユーザ情報
        loginUsers: [],
    },

    // 算出プロパティ
    computed: {
        items() {
            return this.loginUsers;
        },
    },

    // インスタンス作成時の処理
    created: function() {
        this.initInputValue();
        this.doFindAllLoginUsers();
    },

    // メソッド定義
    methods: {
        // ユーザ情報を全て取得する
        doFindAllLoginUsers : function() {
            axios.get('/findAllLoginUsers')
            .then(response => {
                if (response.status != 200) {
                    throw new Error('レスポンスエラー')
                } 
                else {
                    // サーバから取得したログインユーザ情報を設定する
                    //this.loginUsers = response.data.map(x => ({ ...x, userId: x.userId != this.userId }));
                    // test
                    this.loginUsers = response.data;    

                    // TODO 送信対象ユーザ一覧を設定
                }
            }).catch(error =>{
                const {
                    status, 
                    statusText,
                } = error.response;
                console.log(`Error! HTTP Status: ${status} ${statusText}`);
            })
        },
        // 入力値を初期化する
        initInputValue() {
        }
    }
})