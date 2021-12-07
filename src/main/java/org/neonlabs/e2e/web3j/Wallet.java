// package org.neonlabs.e2e.web3j;

// import java.math.BigInteger;
// import java.security.InvalidAlgorithmParameterException;
// import java.security.NoSuchAlgorithmException;
// import java.security.NoSuchProviderException;
// import org.web3j.crypto.CipherException;
// import org.web3j.crypto.ECKeyPair;
// import org.web3j.crypto.Keys;
// import org.web3j.crypto.WalletFile;

// /**
//  * Creates a new wallet.
//  */
// public class Wallet {
//   public static void process(String seed) {

//   // JSONObject processJson = new JSONObject();

//   try {
//   ECKeyPair ecKeyPair = Keys.createEcKeyPair();
//   BigInteger privateKeyInDec = ecKeyPair.getPrivateKey();

//   String sPrivatekeyInHex = privateKeyInDec.toString(16);

//   //
//   Wallet.process(seed);
//   //

//   WalletFile aWallet = Wallet.process(seed); //.createLight(seed,
//   ecKeyPair);
//   String sAddress = aWallet.getAddress();

//   // processJson.put("address", "0x" + sAddress);
//   // processJson.put("privatekey", sPrivatekeyInHex);

//   } catch (CipherException e) {
//   //
//   } catch (InvalidAlgorithmParameterException e) {
//   //
//   } catch (NoSuchAlgorithmException e) {
//   //
//   } catch (NoSuchProviderException e) {
//   //
//   }

//   // return processJson;
//   }
// }
