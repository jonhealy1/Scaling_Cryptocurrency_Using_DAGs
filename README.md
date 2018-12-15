
This program creates 4 Blockchains all running in parallel. 3 sidechains and the main chain. Each sidechain forks off from the main chain and is given a certain set of transactions to work on which would be theoretically divided by the address of the person sending a transaction. In this example sidechain A handles transactions sent by account A, sidechain B handles transactions sent by account B, and sidechain C handles transactions sent by account C. After 2 blocks, the sidechains all append their final values to the main chain. <br/>

For a more detailed explanation of how this system works, see the pdf file provided in this repository. Note: In this example all of the accounts start with 100$. Chains are only making changes to values contained on that chain and because each chain is responsible for a different set of accounts sending value, no account can spend more money than they have.<br/>

./Parallel_POW<br/>

1: Account A sends Account B 10$...<br/>
2: Account B sends Account C 10$...<br/>
3: Account C sends Account A 10$...<br/>

814e325659746b0a3e86dcf33e2ffc2977c93669cb000ad31e3128906c0a9ab0  do more work!  3
d3590d7547d4c05a68cc6bd2daea320e51feb0c35937f46df09985a3cdfd15ba  do more work!  1
7081e1eb6986ba05c65b1464b50e10c211a2cc60fbf2431f436c5a5df8833702  do more work!  2
5a74a6a4c0413cccd764188c540ac9d96a8a6bc842314b9000ec676aa75f8b96  do more work!  3
aca5fadede76662c14a14498c5a1a101bc3dc31a2f78acac9b4ab385056a8981  do more work!  1
218ae0666baf660ec24a60b6404798e186985fb9191788ae276ddf00d207855f  do more work!  2
ebda7cc153030adaba9b6ba67208a764d8edeb60d0ad96a16fa06dbca4b6b4e3  do more work!  2
f164857158c471a00d62151a5c0bfde62980caefd557cef73665a102c4af5b73  do more work!  1
e0bc69239e473546559d3f02c326d94f20e83ac0e02d97e83c5636e8294d1fec  do more work!  3
5bc0036452905284048d96b25b84ace7d6811a4c2b3fb421f25454de7bf9e5bd  do more work!  2
12d7c9018b4e8c4398843ce048158787d19321735d4d6e05019da1a39c774986  do more work!  3
94635dffef16fa1a701befc189626f9f75681fe5a41c116afa2286d90b3acc5a  do more work!  1
d95c86adf10c84cfd173ccd0f3252e838755cf1ad5e07fead81e139f0b1b4f9a  do more work!  1
a89536fb51308603eccbacdcfc0b6c39aa681329255d26bec5dd6739bde87da7  do more work!  2
1fb48b22a7ac7f682f761e706ab04315534e3387483ff85520c895c1c4df7a9f  do more work!  3
4909f24ed94dd2a3b4fb47513a384a47cba0725aebb7b9165eb928d3c2d4a4f0  do more work!  3
b39ba841e78804edb6883cb9bbe251cffcb72dece029fc56ac96b86dda035c00  do more work!  1
26d5e25a3ba7a21e534755fcfaf1cab3743818a77e864081f87155c8c3f1957f  do more work!  2
86aa368a430e93e6d0b1d56b0b8ce21b10f6bc6dd81a8a5683dcc62b67bce0ae  do more work!  2
92c23c978e2882b58ee8a9ea67f4cbd05e0bd81a5860987c0dc7c5a5c7486208  do more work!  3
d256fd0e62ba602b76cffd4b2faec35b8ed65ce97297a1738e7c8f24152f5ee8  do more work!  1
366c7d39413d3da0ac5d513ec1ce280c398f43ba59ab5e05ee3069d2df48815c  do more work!  1
f8baeacb2cf7422adc7c57a634fa9933379634877c35d0b954c079c0abde3766  do more work!  2
8e05fb2a1180cefa1b9b21ac318242415413fbf4c79c1470437c9f830166ed51  do more work!  3
b0fd6fed0b587ead3eb9f237a7de9a1259d5a664f45bfddc93826ba3ae02104c  do more work!  3
6ac7e5a7fd1b327a1e18f9fa16fa27a098492fe3967e242df34231f5a694904b  do more work!  2
77385e8ee5451ec7f987d5e07c7205098f09dbbd64bfd5a17ba7fd0e7a828534  do more work!  1
a6adcefb3c38c35b6ed05ae780fe6f31d639dab1052762cb4d67f2c01fd2b70a  do more work!  1
5279c2554f7982a42bd71b9ff48c79e3c358aa9e0eade9fd18e373cf9149f03f  do more work!  2
27ce17a48ac80b9b604b53ffd9f246d65490a95878bee1e00bc16dae489ea158  do more work!  3
dcc8dfa82704140e9225b7c69a43316b79e2592200b1beabb4550cbb79b9b38d  do more work!  3
a30f1d5e575423ea079698df3897c3ae8ed11a1a968ea5533cdd17e9fdff433f  do more work!  1
7a42d65fb0b6e9b85f99de5e95fee5ef55adcc309e3006e2dbae011a4587374e  do more work!  2
d1b39a38e270071705d90cd25a77a725950cd0d901ec86a7ada1d08c31beeede  do more work!  2
4d1acfc9b7bfeefff26ec4deb11a9e97e22f1cf800951e699bdf14405e520ea3  do more work!  3
5e818eeb84e3d4eb23dad6d763d8e4062511c5263392ce76ba0ada21d21b8816  do more work!  1
800a50f9c5cd465bb947811b51fbb66374a923a6ee04b4481271b8170b39262e  do more work!  2
adb81205d1b6ec7ea810675d126db5168138a14fe89473e0b932a7c40b74f0d0  do more work!  1
344e92c5ac55e7511a8c6b5cc3f981219ae76842602bada86eda14459d246c56  do more work!  3
ce8b129fd70364f5b57bf5c573fe863b78884f054329190f89d85b9efa3711dc  do more work!  3
2ecb9b0f22bab58199d7b9b61874c53e6243a1e33a129f50828ab0149bda90b1  do more work!  2
cfd7e4e1252dea1bb3fefe3c5f4f5c609ee0836d34d910f7742af6bdc97c95ca  do more work!  1
577e580e5dfd05335b0d38dd6dee0bc56e1174dcbc5d9fdd5d0a4361b2a868d1  do more work!  1
f3eec5abbccc193831472ad4c63722714e671c7f206c777cd5f38869d69a15b0  do more work!  3
e49aef95e011673b78adc4630753ca8c4119f554b68360594b8f5704e6ee6e07  do more work!  2
932f66afbc2bac261f1bacd28748f1590fde876f16e691a99398feb30084587f  do more work!  2
e91b981bbb64ed2c42353539d1cea98d65fc6acf03c47aba855eb36733161567  do more work!  1
49642e76d53c0b365776617942201dd1e25db0cdea79308cfde95b4641dfdf78  do more work!  3
93e145d447f56072c6bbd1edfef047f66051a1c5958edc609fcdb04f8ad27d5a  do more work!  2
a01bbdb653a86f01a2d7b7486e21e41ef3024fe9f0ad6eb1ea23996b7817f920  do more work!  3
f8abe9f620dcf5ed3db27b05086157cf1836e61a146fc7c6d33521502f95fcc5  do more work!  1
36ac80fc544a9b7b61b359dee80eb61059378e8c50ed0a1187128298319ac4fc  do more work!  1
57ab53b762aa392b3348de00e02db51a290c45e8c88816d2a2c0f753cf0ad9e2  do more work!  2
7db6a002d43be47ef83a95fb9fe7b668b3bc1ffa3659d482a38dbac6e960ba38  do more work!  3
8a9932676a66ecb4e9fb5241bfa1232f753eb3b500c27cf37c594b084d2a2bec  do more work!  3
bb9909e915c83152226d253dadbff95045c9fde9b323fb773b72a4dcb7519680  do more work!  1
25f00b5f027a3826af07fc87e2aa1e170192667677c8b5beb067a830f8c7dc43  do more work!  2
438d646f177562d260932260cad3ad6a6f8f47883a4abeaf23e302f8cc9f4f7b  do more work!  2
514841db372ddd915ac4a9047eff4295b563b9dc5d0f52b14967c6b2796edf21  do more work!  3
921ce67e4b3f126edd6febac21189edb27f8a42823993d814e4514413fb078c1  do more work!  1
4fe98f6d1fa20ceef45bd18c01ac5c74c990ad1c6434ff29fc899babf49bb9c3  do more work!  1
9a504ccf19863330ea864bc90f3e59d1e9c138c79a89d38f1c89e9ab643475c8  do more work!  3
94063d2388305334abf959b415d756656069094aca4d91c361997fd02bafc536  do more work!  2
d491e25c7d819de5c4e4249603d2e025e93cc2b46f40aa35365dd4ba8eee2dc7  do more work!  3
3eaa41d20914c44dd118dbba46719d78a3aa69f9c8bc0f4c88fbe4c8843d42ff  do more work!  1
306978acba5ad132bb655d98092ec0962eff7bb950339e88ad1542abb350b794  do more work!  2
5e8d5081b19f8d07cc04424d0f64b51ea8a9d4715c3de5b6dbcf155d4de48271  do more work!  2
a20f2f697cbd0017b35193df1634d6f59dfe1592b66c82184bb23f718e3a65cf  do more work!  3
51e87602ab224467289f4e24db75719930ffcb3d088542968fb07c29d0fcc5be  do more work!  1
29366f64d2a8e715f709cfe76d2f8f214e3ec3bef8db86e66bf0af86a6b654b7  do more work!  1
89a934d368457e295242639273b8e98bdf094ef9832ab0693b08f092b81466cd  do more work!  2
9fa9636002270380c17c4d681cbfc814d2424d6a8160c87160e99b4a8fdcea8d  do more work!  3
cc59406a5024b2b2260e0bbfa174d8baf0b905ab3b45ba07f58f327fdd8115c8  do more work!  1
3b76aedbe79bc3da5651aa4f2d83c477fa629c3908df009335328693f3a9a47b  do more work!  3
cfa80cc47552ef23b80e1a4a276ed9fad24688b33be44d73294dcf1aedcf87c0  do more work!  2
b867cb380887f318cdc3222d4ed332d5d08301fbc9d2ec57b31f806c3e2b1f39  do more work!  2
### 040a0c1a4a44afe1ab6ece6e7f04fcdc5ea76b7f8cdfe9d6b9d9c6df5ce8a90d  work done!  A sends B 10
f99d9cb1f4840a6a075233e97639912ebf3f9ecb9ad82a1520ffdaf726096881  do more work!  3
e0721d8a7c40b87c3001b717975ba7195c421930d5997828262b5099bc5cc3a0  do more work!  2
e8a8c00877b040399b844ae3960a6cc51f9212668d840f0fb15408bc2ef7cfc0  do more work!  3
1ded0d557685b7d40974a41dc9c41ed0f0ac1856eecf47e9d8d1c6d9fd41c12e  do more work!  3
81c58735aa0e0f27e66092ad0edcb8a6714631884663af4dd924a803250a96fd  do more work!  2
605c8dfe029843c990b6026a988dfc912abf67da27fa4366bce9de2b008eca22  do more work!  2
b26dfbedce91d7df3cfb8f9dfbfd8aaef763dedeea918b4804366fd5afaae31e  do more work!  3
2069d343de5e12b89388bb7614e943ee359cf0b4447662d86fb597a18d16b862  do more work!  3
e95e233dd06d1788561bdde5f9be1acdc1b1d74f59838e61e0e844513f049ee5  do more work!  2
cb22a583b9cbb613c08a60f7aa0cdb3c77b3dee6ff796eec5b2105620ef7a421  do more work!  2
427579e97fe6298502b95fb68c87f741d4698f47b6e2b3e9b21429072fa1789d  do more work!  3
930199a2e654dba7093c0e3a3d199c888bacf2c1a7a732f930a74bad876d466a  do more work!  2
5f3299985579ce41d99b90999a02ddacdbbccb12ddb5478055886e399b006bd9  do more work!  3
dc11c4f22aaac29facc602903443b2dbbe4dba6ed05b332c718615f38d7b240f  do more work!  3
af3f2148a866fdb1ae4aacc26665ac90e20506ca5a7634377fc6b4e777215882  do more work!  2
6eb65b3d3aae48ce4e7f5f5b55d437a65bf92a70ab7abb3002b3b38f20a6db31  do more work!  2
b675b43ef262cc464363e85437085d48cf1eb6692bc884841f79d60c782e59f9  do more work!  3
6c73272fe49a31480151f5c6a1219f5246cb069674eaf4b25ae3416b4e97b3f2  do more work!  2
bb9e71dbfd4d07d9ba80fd0e23aae65c104fa3ad359896bd2721e1e4d49074fe  do more work!  3
d2d78b7448b9d72f5b2c33d69bc68048f906cd05b044e0789d3514e6de388fb9  do more work!  3
4c3ded13920f2f6c12a7f888812691e30882f2529ee0068e9321b279171893ef  do more work!  2
c12e301a753fd5741db22fd28af31ea60f852915d85cc69638d8b6d1b3eb3b16  do more work!  2
bc5021c3b3118fb7f423a50cb2eed956ba63882be7b66f02ff3c317d2730f670  do more work!  3
4c6b8e166fb3b2214080206dd357b982df8795e3c6e73eac6f48ee30e0345602  do more work!  3
a7bdd44cec9b652118b0c98f10ae1b3d6f640284f62f9f28579f1882d164bb76  do more work!  2
d5694a6db9fc83bd2b546c47d9f99dac9edc9eabf8e49d5b43b046d66a2e3723  do more work!  2
ce378e1a49f8f8bdc8015e82d308bf5ff75ac41e0fb4b929c29d8f8d182782ea  do more work!  3
d3b105d4a22f6d828a1199656341c49633ccb90402f45be5829c4220c4518e7f  do more work!  3
b06d15484968fd61b65e6fcae170f254e19043a461bb542423175e544755ad10  do more work!  2
d3a4ec856b2486301c556cd2dda38ac37b4eca2c4b79c9b2f94832c2a796b367  do more work!  2
103de12e7f8b895c7ed0683fc476e4964b430b16a7c090a41fcd15a693f59970  do more work!  3
dad3392d1133bbd0c5405c4507921d36311423e241d803fc93ccd7132cb187f9  do more work!  3
### 067cb173668c0ff09b1120c36ea3a173e18750736b29ba5c8f39153d6a0c83e6  work done!  B sends C 10
a0de20ccbffd8c226e27c58075333149707a99b390a4bf85da1c3c1d1b4b8475  do more work!  3
b10772d5aebd40a9fce468a5e6585021b848a5b64d1c3d1969b1dc38b99c7e2e  do more work!  3
1e16c23d36ab3b19fe1ecbeef79745aa8289f91ee6875c3b7945921c4c038ec4  do more work!  3
4e2e16a0a87c15151f7e4ea017751bbd493030624941527231e453c49d2022b7  do more work!  3
6b38887662c3e9330fb7ccf2669aab02f3afbac9e5981b50f7e6c9b4f045ad32  do more work!  3
ec120b9635e2eacadd33be0ae402f251f73639a22eb2add7186f329e58465cc9  do more work!  3
ee980525ac4d2af3fbc9a3849430ca293188661f4d825a3e20edd2ea5d7cabac  do more work!  3
333885293db60815fa3ca712372c64f3bb91ff2ac843849345f83f69b1e9eca4  do more work!  3
### 086ea6415d03926edc2a71d9c221a06b05cb261603006219612f4f71eb0a2855  work done!  3: C sends A 10<br/>

4: Account A sends Account C 5$...<br/>
5: Account B sends Account A 5$...<br/>
6: Account C sends Account A 20$...<br/>

4575974d77b7f081cfdfe37c8bc356390a2621bd82a0fe8272841be2362015e5  do more work!  6
9852a8b6070dd32f9692b9dd6d1c530e8100b8c338db4321998ed7b06005c7ed  do more work!  4
e8cf98c12fb67c9a95afaca6b2bb0ea26e5b27387c2b85213b64be3b8ed0628d  do more work!  5
cdc98efa7b2d268cd59297dac087302e18736c1f241d9625575bae1c07803164  do more work!  4
be6844db3eb40b13f94d21f486535c15465a2dc230d74351569cd95f04584513  do more work!  5
23bbbe339105aa4fce8ab32f68756dd6f8024fc08b92bf9e55ff73b76f8746a8  do more work!  6
a709c674a088429ef72982fd6d97959f8526b65ea634a408c242a3cd6d5153ea  do more work!  4
c2f1c40e707b5b605b06f63685b2274fa1a2d4f0d7e5bf99cf3bddacc7aad45d  do more work!  5
b02e56fe10acb4a4c442efb3000fc9e1e1f1c901bc903d707a9b5be07fb17395  do more work!  6
39750f63a8c9964a9a3d974806dd635a33d54bbb85dce87d46296403c4013496  do more work!  6
dc62c182c0a6027f1e0fe1e0ed74e12784855d1aed2428b30300cbd113b15670  do more work!  4
e98d3365aee37ffeab3810ad9da79733679e921016a2545c08a7c512b77eeecc  do more work!  5
4ae3a98d96f94631b7cac778969bc68cfd4217d69ed179e270d4d9f9eb598289  do more work!  5
3e0e0dc29a53326f6e42e6d351051aaef607c21ee52688633011112b0b2844b6  do more work!  4
ca64176a7a49974f72242982a1a87ec63594e770f17eac7d184b8de45e031e97  do more work!  6
8bf640be2d4ce954f293a8051da79f5281c58c93d59070341e545117607ee869  do more work!  6
eef8aa70ae6bb69a7201044bfdaaaf684317505592ccd9dd5675ccd9a3cb63aa  do more work!  5
67439c806edc7553bb3f909785fc80469d220a791bb532304bb7eefc8e7dd93a  do more work!  4
6e52f2566114078f66c9d12cfb10597265692d97733d796a0f6772f01c0bc359  do more work!  4
5cc2a111f38c387e1d82e8b20deb39bc8879298c2023d97228e20a934d0f4cbf  do more work!  6
cac03bbedc8a67ee1c11f8907b70737fbf18b92aa1000221b48011cc96bc881c  do more work!  5
a74a1f3bb7b52c5fdc691f3a6cc3f56d55767fd59f87be255c43bd328c00f120  do more work!  5
5bead5b882f14187c1984f5a7498aa02f768cb6fe7a69959b8e275a34e35fe1b  do more work!  4
3a7b5b3bbe5af4e2c24a130a6f58b885ebf74faf926ed4acabe33fd88984b32c  do more work!  6
### 08d725d43e75f637c582c41b0d15334f33a6e27f858b7ac73774b13dcf7ea21b  work done!  C sends A 20
8400e06b3eb50d9c6cb7c045a7dbe973d2f3d768a0bed978a03b57b63b759294  do more work!  5
620133abd45dd794f2452c916c9bf0a88ce5446d412b2f8ff2ee34617ee98b8a  do more work!  4
d6efa19a5024927e227acb9a794eb80fd648ee1fd5266d240e6fad7edfe331af  do more work!  5
cef0b833eb6783cc132280fb20aa00ed04126fd9a565b7be157cf60ce2291110  do more work!  4
9614558b7feb13cd5a16196d45b2f1c19b2f7f8eb7384e1988eb2d858e82a1b6  do more work!  4
ff483691f9762353e95d3339f16d7aef38750ccd28f9b30dc08a06a263489aa0  do more work!  5
156cb6a0bfab1ed6ec010c6d4f1f7e72909c242e974d55fc95821958cc1d6cb5  do more work!  4
edc02d7463225ea65c84ef20e6924d3146e42d253551ec50aeace1f84207128a  do more work!  5
### 0a01a59100ff0d32241ad136698ee1f3d1dda4cc47ce6946d4aca01893c8a2fd  work done!  B sends A 5
2f1f76dd09c5983a274a63e0cf0930ad80be58533d79807034c7f7825aa4747a  do more work!  4
2c7618baf1285e9db4a126cf368263bc415e5145af76d9ae4d5296b0529cadf9  do more work!  4
bcfc8b0afaeb43dbefbc33fcff019e128c631563f4951e0e5f2729facec25223  do more work!  4
b17088e88b9fa1dfacdaa993dd7e42fa2acc8485a4031c96f5d7d1ae7c372695  do more work!  4
c037eba57dc3f52223723d9b87e1904863c2609479dfc126592d4369ca45dd4f  do more work!  4
aad562453a3ab285d3e1183d95a41627ed7bb4d2e74ae6804267429993aacf66  do more work!  4
87f6b797d8916faa463ae0335c467e689795939c3d5a8cf6f62d6299f0bea293  do more work!  4
96a1ea449392324dae4e2191b639842bbba8f0ed33b596cb7588a042789ec27c  do more work!  4
53afe77298c84cc6e6cb173e840c901d057078b3dd70eabab5ad5d8fb2ece84f  do more work!  4
dc278d03338569f8cd0f8c42d5f6296ac69d6463605554ac3f865bde102b8d12  do more work!  4
b278af1507a9f2713282371fb99978e9756605bb9c6d17f91077d3db6437eaa0  do more work!  4
7961fdf1fea39e2d62a855f1364b41eaa438eb28d6cdd1f7aed1cc7ce2776e0f  do more work!  4
8184ca5e7bc282be34c665f13745ceeec6df29e7a2879c52704f7141452e04ff  do more work!  4
a56d6348d11bd435399130b9981f7b58a2d68d50e9715dca10ba5ae5b21e9ff9  do more work!  4
c3862e53532085341c880f0336d545316ab5f752c4b3ea0ef2c768e51f620e1e  do more work!  4
46e6b7d09c91ce34a004a95aed603ca60afa823546ed1d3e977361bdb99cb0a0  do more work!  4
b00975b566449fa99c76e4e9cd087a03a96b387849fe276ca722e9ba448f3121  do more work!  4
c9cbe1dae68fa0dcf476cc24ac8020f5d1e6ac863970c35b5d31b1ad32d2ecfa  do more work!  4
fbd68eaaa242d5db1c6ae81ae6c56c2f5dd6fa5aaf6b56ccf9a28a51ca3a3499  do more work!  4
27db6753d03de8cb0a3282fd31cb5befb8ffcb3aee00679a41229260376a8247  do more work!  4
b3fa86c90f41aa936ff683392ee494c77b3336a06f097c07964315cd5e6aeeb4  do more work!  4
b550e9abac215af9224d94a0342a0d2b2a799f467163f26f3a4bbe5ffdce2463  do more work!  4
### 015c31186d6fef67c81960f370f35f0df8c9e388059eee2e8f636584322af359  work done!  A sends C 5<br/>

Side Chain A, Block: 0 {0 2018-12-13 14:58:54.624422058 -0800 PST 100 100 100 9af15b336e6a9619928537df30b2e6a2376569fcf9d7e773eccede65606529a0  1 }<br/>

Side Chain A, Block: 1 {1 2018-12-13 14:58:54.625475699 -0800 PST 90 110 100 040a0c1a4a44afe1ab6ece6e7f04fcdc5ea76b7f8cdfe9d6b9d9c6df5ce8a90d 9af15b336e6a9619928537df30b2e6a2376569fcf9d7e773eccede65606529a0 1 19}<br/>

Side Chain A, Block: 2 {2 2018-12-13 14:59:44.710988131 -0800 PST 85 110 105 015c31186d6fef67c81960f370f35f0df8c9e388059eee2e8f636584322af359 040a0c1a4a44afe1ab6ece6e7f04fcdc5ea76b7f8cdfe9d6b9d9c6df5ce8a90d 1 22}<br/>

Side Chain B, Block: 0 {0 2018-12-13 14:58:54.624422058 -0800 PST 100 100 100 9af15b336e6a9619928537df30b2e6a2376569fcf9d7e773eccede65606529a0  1 }<br/>

Side Chain B, Block: 1 {1 2018-12-13 14:58:54.625556299 -0800 PST 100 90 110 067cb173668c0ff09b1120c36ea3a173e18750736b29ba5c8f39153d6a0c83e6 9af15b336e6a9619928537df30b2e6a2376569fcf9d7e773eccede65606529a0 1 29}<br/>

Side Chain B, Block: 2 {2 2018-12-13 14:59:44.711069385 -0800 PST 105 85 110 0a01a59100ff0d32241ad136698ee1f3d1dda4cc47ce6946d4aca01893c8a2fd 067cb173668c0ff09b1120c36ea3a173e18750736b29ba5c8f39153d6a0c83e6 1 c}<br/>   

Side Chain C, Block: 0 {0 2018-12-13 14:58:54.624422058 -0800 PST 100 100 100 9af15b336e6a9619928537df30b2e6a2376569fcf9d7e773eccede65606529a0  1 }<br/>

Side Chain C, Block: 1 {1 2018-12-13 14:58:54.625470568 -0800 PST 110 100 90 086ea6415d03926edc2a71d9c221a06b05cb261603006219612f4f71eb0a2855 9af15b336e6a9619928537df30b2e6a2376569fcf9d7e773eccede65606529a0 1 32}<br/>

Side Chain C, Block: 2 {2 2018-12-13 14:59:44.710979405 -0800 PST 130 100 70 08d725d43e75f637c582c41b0d15334f33a6e27f858b7ac73774b13dcf7ea21b 086ea6415d03926edc2a71d9c221a06b05cb261603006219612f4f71eb0a2855 1 8}<br/>

Main Block Chain, Block: 0 {0 2018-12-13 14:58:54.624422058 -0800 PST 100 100 100 9af15b336e6a9619928537df30b2e6a2376569fcf9d7e773eccede65606529a0  1 }<br/>

Main Block Chain, Block: 1 {1 2018-12-13 14:58:54.625475699 -0800 PST 90 110 100 040a0c1a4a44afe1ab6ece6e7f04fcdc5ea76b7f8cdfe9d6b9d9c6df5ce8a90d 9af15b336e6a9619928537df30b2e6a2376569fcf9d7e773eccede65606529a0 1 19}<br/>

Main Block Chain, Block: 2 {2 2018-12-13 14:59:44.710988131 -0800 PST 85 110 105 015c31186d6fef67c81960f370f35f0df8c9e388059eee2e8f636584322af359 040a0c1a4a44afe1ab6ece6e7f04fcdc5ea76b7f8cdfe9d6b9d9c6df5ce8a90d 1 22}<br/>

Main Block Chain, Block: 3 {1 2018-12-13 14:58:54.625556299 -0800 PST 100 90 110 067cb173668c0ff09b1120c36ea3a173e18750736b29ba5c8f39153d6a0c83e6 9af15b336e6a9619928537df30b2e6a2376569fcf9d7e773eccede65606529a0 1 29}<br/>

Main Block Chain, Block: 4 {2 2018-12-13 14:59:44.711069385 -0800 PST 105 85 110 0a01a59100ff0d32241ad136698ee1f3d1dda4cc47ce6946d4aca01893c8a2fd 067cb173668c0ff09b1120c36ea3a173e18750736b29ba5c8f39153d6a0c83e6 1 c}<br/>

Main Block Chain, Block: 5 {1 2018-12-13 14:58:54.625470568 -0800 PST 110 100 90 086ea6415d03926edc2a71d9c221a06b05cb261603006219612f4f71eb0a2855 9af15b336e6a9619928537df30b2e6a2376569fcf9d7e773eccede65606529a0 1 32}<br/>

Main Block Chain, Block: 6 {2 2018-12-13 14:59:44.710979405 -0800 PST 120 95 85 08d725d43e75f637c582c41b0d15334f33a6e27f858b7ac73774b13dcf7ea21b 086ea6415d03926edc2a71d9c221a06b05cb261603006219612f4f71eb0a2855 1 8}<br/>

*** Account Values on Each Sidechain ***<br/>
* Only accounts that a sidechain is responsible for get decremented<br/>
Sidechain A:  {85 110 105}<br/>
Sidechain B:  {105 85 110}<br/>
Sidechain C:  {130 100 70}<br/>

*** Final Results on the Main Blockchain ***<br/>
* These are the average results for the accounts taken from the values recorded on each sidechain<br/>
* For example Account A spent 15 in total on sidechain A and received 5 on B and 30 on C, for a grand total of 20. Each account started with 100 so A has 100 + 20 at the end for 120. <br/>
A: 120<br/>
B: 95<br/>
C: 85<br/>
<br/>
 * Future Work *<br/>
In the future I would like finish all of the consensus rules and then deploy the more developed version of this implementation online for some serious testing. This is really just meant to be a parallel version of Bitcoin but there are still a lot of things to iron out with having sidechains and deciding how to delegate miners primarily looking to keep the system balanced and secure. Updated rules for validator nodes and forking more chains under load are other things to think about as well. I think the benefit of this system potentially could be in its simplicity.<br/>
